package main

import (
	"errors"
	"fmt"

	"github.com/jdetok/golib/geterr"
)

func main() {
	err := testErr()
	if err != nil {
		fmt.Println(err)
	}
}

func testErr() error {
	e := geterr.InitErr()
	e.Msg = "An error occured"
	return e.BuildErr(errors.New("error occured"))
}
