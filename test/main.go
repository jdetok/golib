package main

import (
	"fmt"

	"github.com/jdetok/golib/logdeko"
	"github.com/jdetok/golib/maild"
)

func main() {
	l, err := logdeko.InitLogF("testd", "testf")
	if err != nil {
		fmt.Println(err)
	}
	l.WriteLog("writing to log testing email")

	m := maild.MakeMail(
		[]string{"jdeko17@gmail.com", "jdekock17@gmail.com"},
		"Test Subject from Main",
		"Testing the body!!!!",
	)

	if err := m.SendBasicEmail(); err != nil {
		fmt.Println(err)
	}

	// m.Attach(l.LogF)
	// fmt.Println(m.File)

	// err := maild.SendBasicEmail("another test subject",
	// 	"NEW STYLE!this is my test body")
	// if err != nil {
	// 	fmt.Println(err)
	// }
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
