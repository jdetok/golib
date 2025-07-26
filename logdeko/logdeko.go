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

func InitLogF(dir string, file string) (Logger, error) {
	e := geterr.InitErr()
	var logd = Logger{
		Dir:  dir,
		File: file,
	}
	if err := logd.MakeLogF(); err != nil {
		e.Msg = "failed to init logger"
		return Logger{}, e.BuildErr(err)
	}
	return logd, nil
}

func (l *Logger) WriteLog(msg string) error {
	e := geterr.InitErr()

	// open file to write (append)
	f, err := os.OpenFile(l.LogF, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		e.Msg = fmt.Sprintf("failed to open %s", l.LogF)
	}

	// pass f & msg to Fprintf to write msg string followed by line break to the file
	n, err := fmt.Fprintf(f, "-- %s\n", msg)
	if err != nil {
		e.Msg = "error writing to log file"
		return e.BuildErr(err)
	}
	fmt.Printf("wrote %d bytes to %s\n", n, l.LogF)
	return nil
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

func (l *Logger) BuildPath() string {
	return fmt.Sprintf("%s/%s_%s.log",
		l.Dir, l.File, time.Now().Format("010206_150405"))
}
