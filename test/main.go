package main

import (
	"fmt"
	"time"

	"github.com/jdetok/golib/logdeko"
)

func main() {
	var logd = logdeko.Logger{
		Dir:  "./testlog",
		File: "test",
	}
	if err := logd.MakeLogF(); err != nil {
		fmt.Println(err)
	}
	err := logd.WriteLog(fmt.Sprintf("Hello it is %v", time.Now()))
	if err != nil {
		fmt.Println(err)
	}
}
