package maild

import (
	"fmt"
	"net/smtp"

	"github.com/jdetok/golib/getenv"
	"github.com/jdetok/golib/geterr"
)

type MIMEmail struct {
	User string
	Pass string
	Host string
	Port string
	Addr string
	Subj string
	Body string
	Mesg string
	File string
	MlTo []string
	Auth smtp.Auth
}

// constructor
func MakeMail(mailTo []string, subject, body string) MIMEmail {
	var m MIMEmail
	getenv.LoadDotEnv()
	m.User = getenv.EnvStr("GMAIL_SNDR")
	m.Pass = getenv.EnvStr("GMAIL_PASS")
	m.Host = getenv.EnvStr("GMAIL_HOST")
	m.Port = getenv.EnvStr("GMAIL_PORT")
	m.MlTo = mailTo
	m.Subj = subject
	m.Body = body
	m.MakeAddr()
	return m
}

func (m *MIMEmail) MakeAddr() {
	m.Addr = fmt.Sprint(m.Host, ":", m.Port)
}

func (m *MIMEmail) MakeBasicEmail() {
	m.Mesg = fmt.Sprintf("Subject: %s\n%s", m.Subj, m.Body)
}

func (m *MIMEmail) AuthGmail() {
	m.Auth = smtp.PlainAuth("", m.User, m.Pass, m.Host)
}

func (m *MIMEmail) SendBasicEmail() error {
	e := geterr.InitErr()
	m.AuthGmail()
	m.MakeBasicEmail()
	err := smtp.SendMail(m.Addr, m.Auth, m.User, m.MlTo, []byte(m.Mesg))
	if err != nil {
		e.Msg = "error sending email"
		return e.BuildErr(err)
	}
	return nil
}

func (m *MIMEmail) Attach(fName string) {
	m.File = fName
}

/*
func AuthGmail() smtp.Auth {
	getenv.LoadDotEnv()
	user := getenv.EnvStr("GMAIL_SNDR")
	pw := getenv.EnvStr("GMAIL_PASS")
	addr := getenv.EnvStr("GMAIL_HOST")
	return smtp.PlainAuth("", user, pw, addr)
}


func SendBasicEmail(subject, body string) error {
	e := geterr.InitErr()
	auth := AuthGmail()

	msg := MakeEmail(subject, body)
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
*/
