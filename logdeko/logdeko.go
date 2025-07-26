package logdeko

import (
	"fmt"
	"os"
	"time"

	"github.com/jdetok/golib/geterr"
)

type Logger struct {
	Dir  string
	File string
	LogF string
}

func (l *Logger) BuildPath() string {
	return fmt.Sprintf("%s/%s_%s.log",
		l.Dir, l.File, time.Now().Format("010206_150405"))
}

func (l *Logger) MakeLogF() error {
	e := geterr.InitErr()

	// create directory if it doesn't exist
	if err := os.MkdirAll(l.Dir, 0750); err != nil {
		e.Msg = fmt.Sprintf("failed to create directory %s", l.Dir)
		return e.BuildErr(err)
	}

	// create file
	l.LogF = l.BuildPath()
	f, err := os.Create(l.LogF)
	if err != nil {
		e.Msg = fmt.Sprintf("failed to create file %s", l.LogF)
		return e.BuildErr(err)
	}
	defer f.Close()
	return nil
}

func (l *Logger) WriteLog(msg string) error {
	e := geterr.InitErr()
	f, err := os.OpenFile(l.LogF, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		e.Msg = fmt.Sprintf("failed to open %s", l.LogF)
	}
	n, err := f.WriteString(msg)
	if err != nil {
		e.Msg = "error writing to log file"
		return e.BuildErr(err)
	}
	fmt.Printf("wrote %d bytes to %s\n", n, l.LogF)
	return nil
}
