package main

import (
	"fmt"

	"github.com/jdetok/golib/logd"
	"github.com/jdetok/golib/maild"
)

func main() {
	l, err := logd.InitLogF("testd", "testf")
	if err != nil {
		fmt.Println(err)
	}
	l.WriteLog("writing to log testing email")

	m := maild.MakeMail(
		[]string{"jdeko17@gmail.com", "jdekock17@gmail.com"},
		"Test Subject from Main",
		"Testing the body!!!!",
	)

	if err := m.SendMIMEEmail(l.LogF); err != nil {
		fmt.Println(err)
	}
}
