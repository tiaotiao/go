package util

import (
	"fmt"
	"reflect"
)

// Find the positon of val in the array. Return -1 if does not found.
func Find(array interface{}, val interface{}) int {
	v := valueOfSlice(array)

	if v.Type().Elem() != reflect.TypeOf(val) {
		panic("Elem and Val type not match")
	}

	for i := 0; i < v.Len(); i++ {
		ei := v.Index(i)
		if reflect.DeepEqual(ei.Interface(), val) {
			return i
		}
	}

	return -1
}

// Join all elements of array as a string seperated by sep.
func Join(array interface{}, sep string) string {
	v := valueOfSlice(array)

	s := ""
	n := v.Len()
	for i := 0; i < n; i++ {
		if i > 0 {
			s += sep
		}
		e := v.Index(i)
		s += fmt.Sprintf("%v", e.Interface())
	}
	return s
}

// Reverse elements of array from beinning to end.
func Reverse(array interface{}) {
	v := valueOfSlice(array)

	var tmp reflect.Value = reflect.New(v.Type().Elem()).Elem()

	for i, j := 0, v.Len()-1; i < j; i, j = i+1, j-1 {
		ei := v.Index(i)
		ej := v.Index(j)
		tmp.Set(ei)
		ei.Set(ej)
		ej.Set(tmp)
	}
}

// Insert an element into array at the position of index.
func Insert(ptrArray interface{}, index int, val interface{}) {
	v := valueOfSlicePtr(ptrArray)
	if v.Type().Elem() != reflect.TypeOf(val) {
		panic("type not match")
	}
	if index > v.Len() {
		panic("Insert index out of bounds")
	}

	a := v.Slice(0, index)
	b := v.Slice(index, v.Len())

	c := v
	if v.Len() == v.Cap() {
		// extend capacity
		ext := v.Cap() / 2
		if ext < 8 {
			ext = 8
		}
		c = reflect.MakeSlice(v.Type(), v.Len()+1, v.Cap()+ext)
	} else {
		c.SetLen(v.Len() + 1)
	}

	reflect.Copy(c, a)
	reflect.Copy(c.Slice(index+1, c.Len()), b)

	c.Index(index).Set(reflect.ValueOf(val))
	v.Set(c)
}

// Delete all elems equals to Val in Array.
// Return the number of elems removed.
func Delete(ptrArray interface{}, val interface{}) int {
	check := func(e interface{}) bool {
		return reflect.DeepEqual(e, val)
	}
	return DeleteEx(ptrArray, check)
}

// Delete all elems which check() returns true.
// Return the number of elems removed.
func DeleteEx(ptrArray interface{}, check func(e interface{}) bool) int {
	v := valueOfSlicePtr(ptrArray)

	var removed int
	for i, j := 0, 0; i < v.Len(); i++ {
		ei := v.Index(i)
		if check(ei.Interface()) {
			removed += 1
		} else {
			v.Index(j).Set(ei)
			j += 1
		}
	}

	v.SetLen(v.Len() - removed)

	return removed
}

func valueOfSlice(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	if t.Kind() != reflect.Slice {
		panic("not slice type")
	}
	return v
}

func valueOfSlicePtr(i interface{}) reflect.Value {
	pv := reflect.ValueOf(i)
	pt := reflect.TypeOf(i)
	if pt.Kind() != reflect.Ptr {
		panic("not a pointer")
	}
	t := pt.Elem()
	v := pv.Elem()
	if t.Kind() != reflect.Slice {
		panic("not slice type")
	}
	if !v.CanAddr() {
		panic("cannot address")
	}
	return v
}
