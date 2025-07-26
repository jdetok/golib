package maild

import (
	"fmt"
	"net/smtp"

	"github.com/jdetok/golib/getenv"
	"github.com/jdetok/golib/geterr"
)

func SendEmail(msg string) error {
	e := geterr.InitErr()
	getenv.LoadDotEnv()
	user := getenv.EnvStr("GMAIL_SNDR")
	pw := getenv.EnvStr("GMAIL_PASS")
	addr := getenv.EnvStr("GMAIL_URL")

	auth := smtp.PlainAuth(
		"",
		user,
		pw,
		addr,
	)

	err := smtp.SendMail(
		fmt.Sprintf(
			"%s:%s", addr, getenv.EnvStr("GMAIL_PORT")),
		auth,
		user,
		[]string{user, "jdeko17@gmail.com", "jdekock17@gmail.com"},
		[]byte(msg),
	)

	if err != nil {
		e.Msg = "error sending email"
		return e.BuildErr(err)
	}
	return nil
}

func MakeEmail(subject string, body string) string {
	return fmt.Sprintf("Subject: %s\n%s", subject, body)
}
