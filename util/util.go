package util

import (
	"fmt"
	"runtime"
)

// Recover panic and print stack.
// Notice: should be used in defer.
func Recover(printStack bool) (err error) {
	e := recover()
	if e == nil {
		return nil
	}

	if v, ok := e.(error); ok {
		err = fmt.Errorf("Panic: %v", v.Error())
	} else {
		err = fmt.Errorf("Panic: %v", e)
	}

	if printStack {
		println(err.Error(), "\n")

		buf := make([]byte, 2000)
		n := runtime.Stack(buf, false)
		println(string(buf[:n]))
	}

	return err
}
