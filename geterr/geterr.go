package geterr

import (
	"fmt"
	"runtime"
)

type Err struct {
	Func string
	Msg  string
}

func (e *Err) BuildErr(err error) error {
	startEnd := "************"
	return fmt.Errorf("%s\n** ERROR OCCURED IN %s\n*** MSG: %s\n*** ERR MSG FROM FUNC: %e\n%s",
		startEnd,
		e.Func,
		e.Msg,
		err,
		startEnd)
}

func InitErr() Err {
	var e Err
	pc, _, _, _ := runtime.Caller(1)
	e.Func = runtime.FuncForPC(pc).Name()
	return e
}
