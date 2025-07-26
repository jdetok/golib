package main

import (
	"fmt"

	maild "github.com/jdetok/golib/mail"
)

func main() {
	err := maild.SendBasicEmail("another test subject",
		"NEW STYLE!this is my test body")
	if err != nil {
		fmt.Println(err)
	}
	/*
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
	*/
}
