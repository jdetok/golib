package main

import (
	"fmt"

	maild "github.com/jdetok/golib/mail"
)

func main() {
	msg := maild.MakeEmail("test subject", "hello this is my test body")
	if err := maild.SendEmail(msg); err != nil {
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
