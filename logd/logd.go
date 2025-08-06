package logd

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/jdetok/golib/errd"
)

type Logger struct {
	Dir  string
	File string
	LogF string
}

type Logd struct {
	Msg string
}

// TODO - write to top of file on init
func InitLogger(dir string, file string) (Logger, error) {
	e := errd.InitErr()
	var l = Logger{
		Dir:  dir,
		File: file,
	}
	if err := l.MakeLogF(); err != nil {
		e.Msg = "failed to init logger"
		return Logger{}, e.BuildErr(err)
	}
	l.LogHead()
	return l, nil
}

func (l *Logger) LogHead() {
	e := errd.InitErr()

	f, err := os.OpenFile(l.LogF, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		e.Msg = fmt.Sprintf("failed to open %s", l.LogF)
		fmt.Println(e.BuildErr(err))
	}
	fmt.Fprintf(f, "***** RUN TIME %v\n\n", time.Now())
}

func (l *Logger) WriteLog(msg string) {
	e := errd.InitErr()

	f, err := os.OpenFile(l.LogF, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		e.Msg = fmt.Sprintf("failed to open %s", l.LogF)
		fmt.Println(e.BuildErr(err))
	}

	// LogFunc will log the name of the function before the message
	var ld Logd
	pc, _, _, _ := runtime.Caller(1)
	ld.Msg = fmt.Sprintln("+++++", runtime.FuncForPC(pc).Name())
	ld.Msg += fmt.Sprintln("--", msg)

	// print to console & write to log file
	fmt.Println(ld.Msg)
	_, err = fmt.Fprintln(f, ld.Msg)
	if err != nil {
		e.Msg = "error writing to log file"
		fmt.Println(e.BuildErr(err))
	}
	// fmt.Printf("wrote %d bytes to %s\n", n, l.LogF)
}

func (l *Logger) MakeLogF() error {
	e := errd.InitErr()

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
