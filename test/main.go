package main

import (
	"fmt"
	"time"

	"github.com/jdetok/golib/logdeko"
)

func main() {
	logd, err := logdeko.InitLogF("testd", "testf")
	if err != nil {
		fmt.Println(err)
	}

	err = logd.WriteLog(fmt.Sprintf("Hello it is %v", time.Now()))
	if err != nil {
		fmt.Println(err)
	}

	err = logd.WriteLog(fmt.Sprintf("Hello again it is now %v", time.Now()))
	if err != nil {
		fmt.Println(err)
	}
}
