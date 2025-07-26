package errd

import (
	"errors"
	"fmt"
	"runtime"
)

type Err struct {
	Func string
	Msg  string
}

func InitErr() Err {
	var e Err
	pc, _, _, _ := runtime.Caller(1)
	e.Func = runtime.FuncForPC(pc).Name()
	return e
}

// to be used when an error comes from a called function with its own err returned
func (e *Err) BuildErr(err error) error {
	startEnd := "************"
	return fmt.Errorf("%s\n** ERROR OCCURED IN %s\n** MSG: %s\n** ERR MSG FROM FUNC: %e\n%s",
		startEnd,
		e.Func,
		e.Msg,
		err,
		startEnd)
}

// to be used when there is no existing error to pass to BuildErr
func (e *Err) NewErr() error {
	startEnd := "************"
	return fmt.Errorf("%s\n** ERROR OCCURED IN %s\n** MSG: %s\n%s",
		startEnd,
		e.Func,
		e.Msg,
		startEnd)
}

func TestErr() error {
	e := InitErr()
	e.Msg = "An error occured"
	return e.BuildErr(errors.New("error occured"))
}
