package maild

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"strings"

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

func (m *MIMEmail) MakeMIMEMsg(fName string) error {
	e := geterr.InitErr()
	atch, err := m.Attach(fName)
	if err != nil {
		e.Msg = fmt.Sprintf("error attaching file at %s", fName)
		return e.BuildErr(err)
	}
	bndry := "bndry-" + fName
	m.Mesg = strings.Join([]string{
		"From: " + m.User,
		"To: " + strings.Join(m.MlTo, ", "),
		"Subject: " + m.Subj,
		"MIME-Version: 1.0",
		"Content-Type: multipart/mixed; boundary=" + bndry,
		"",
		"--" + bndry,
		"Content-Type: text/plain; charset=utf-8",
		"Content-Transfer-Encoding: 7bit",
		"",
		m.Body,
		"",
		"--" + bndry,
		"Content-Type: application/octet-stream",
		fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"", m.File),
		"Content-Transfer-Encoding: base64",
		"",
		atch,
		"--" + bndry + "--",
		"",
	}, "\r\n")
	return nil
}

func (m *MIMEmail) SendMIMEEmail(fName string) error {
	e := geterr.InitErr()

	m.AuthGmail()
	err := m.MakeMIMEMsg(fName)
	if err != nil {
		e.Msg = fmt.Sprintf(
			"failed to create MIME msg with file attached at %s", fName)
		return e.BuildErr(err)
	}

	err = smtp.SendMail(m.Addr, m.Auth, m.User, m.MlTo, []byte(m.Mesg))
	if err != nil {
		e.Msg = "error sending email"
		return e.BuildErr(err)
	}
	return nil
}

func (m *MIMEmail) Attach(fName string) (string, error) {
	e := geterr.InitErr()

	m.File = fName
	// read file at fName as []byte
	f, err := os.ReadFile(m.File)
	if err != nil {
		e.Msg = fmt.Sprintf("error reading file at %s", m.File)
		return "", e.BuildErr(err)
	}

	// encode bytes to base64 string
	enc := base64.StdEncoding.EncodeToString(f)
	return SplitFileLines(76, enc, "\r\n"), nil
}

func SplitFileLines(cLen int, body, delim string) string {
	var b strings.Builder
	for len(body) > cLen {
		b.WriteString(body[:cLen])
		b.WriteString(delim)
		body = body[cLen:]
	}
	b.WriteString(body)
	b.WriteString(delim)
	return b.String()
}
