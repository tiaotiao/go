package reflectutil

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
	pv := reflect.ValueOf(ptr)
	if pv.Kind() != reflect.Ptr {
		panic("ptr must be a pointer")
	}
	ev := pv.Elem()
	if !ev.CanSet() {
		panic("elem of ptr can not be set")
	}

	vv := reflect.ValueOf(val)

	if vv.Type() != ev.Type() {
		return fmt.Errorf("type not match (%v, %v)", vv.Type().String(), ev.Type().String())
	}

	ev.Set(vv)

	return nil
}

func TypeEqual(v1, v2 interface{}) bool {
	t1 := reflect.TypeOf(v1)
	t2 := reflect.TypeOf(v2)
	return t1 == t2
}
