package util

import (
	"fmt"
	"reflect"
)

// Assign val to the address of ptr. And the type of val must be *ptr.
// You can think that it works like this:
//   v, ok := val.(type(*ptr))
//   if !ok { panic() }
//   *ptr = v
func Assign(ptr interface{}, val interface{}) error {
	ev, err := ptrElemOf(ptr)
	if err != nil {
		return err
	}

	if val == nil {
		ev.Set(reflect.Zero(ev.Type()))
		return nil
	}

	vv := reflect.ValueOf(val)

	if !vv.Type().ConvertibleTo(ev.Type()) {
		return fmt.Errorf("type not match (%v, %v) (%v, %v)", ptr, ev.Type().String(), val, vv.Type().String())
	}

	ev.Set(vv)

	return nil
}

func TypeEqual(v1, v2 interface{}) bool {
	t1 := reflect.TypeOf(v1)
	t2 := reflect.TypeOf(v2)
	return t1 == t2
}

func MustPointer(v interface{}) {
	_, err := ptrElemOf(v)
	if err != nil {
		panic(err.Error())
	}
}

func ptrElemOf(v interface{}) (reflect.Value, error) {
	pv := reflect.ValueOf(v)
	if pv.Kind() != reflect.Ptr {
		return pv, fmt.Errorf("'%v(%v)' not a pointer", pv, pv.Kind())
	}
	ev := pv.Elem()
	if !ev.CanSet() {
		return pv, fmt.Errorf("'%v(%v)' can not set", pv, pv.Kind())
	}
	return ev, nil
}
