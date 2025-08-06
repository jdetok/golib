package errd

import (
	"errors"
	"fmt"
	"runtime"
)

const DELIM string = "************"
const EMSG string = "%s\n** ERROR OCCURED IN %s\n** MSG: %s\n** %s ERR MSG: %e\n%s"

type Err struct {
	Func string
	Msg  string
}

// captures name of function that called it
func InitErr() Err {
	var e Err
	pc, _, _, _ := runtime.Caller(1)
	e.Func = runtime.FuncForPC(pc).Name()
	return e
}

// to be used when an error comes from a called function with its own err returned
func (e *Err) BuildErr(err error) error {
	// startEnd := "************"
	return fmt.Errorf(
		// "%s\n** ERROR OCCURED IN %s\n** MSG: %s\n** %s ERR MSG: %e\n%s",
		EMSG,
		DELIM,
		e.Func,
		e.Msg,
		e.Func,
		err,
		DELIM)
}

func (e *Err) BuildErrStr(err error) string {
	// startEnd := "************"
	return fmt.Sprintf(
		// "%s\n** ERROR OCCURED IN %s\n** MSG: %s\n** ERR MSG FROM %s: %e\n%s",
		EMSG,
		DELIM,
		e.Func,
		e.Msg,
		e.Func,
		err,
		DELIM)
}

// to be used when there is no existing error to pass to BuildErr
func (e *Err) NewErr() error {
	return fmt.Errorf("%s\n** ERROR OCCURED IN %s\n** MSG: %s\n%s",
		DELIM,
		e.Func,
		e.Msg,
		DELIM)
}

func TestErr() error {
	e := InitErr()
	e.Msg = "An error occured"
	return e.BuildErr(errors.New("error occured"))
}
