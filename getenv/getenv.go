package getenv

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jdetok/golib/geterr"
	"github.com/joho/godotenv"
)

func LoadDotEnv() error {
	e := geterr.Err{Func: "getenv.LoadDotEnv()"}
	if err := godotenv.Load(); err != nil {
		e.Msg = "*** FATAL: failed to load .env variabels"
		return e.BuildErr(err)
	}
	return nil
}

func GetEnvStr(key string) (string, error) {
	e := geterr.Err{Func: "GetEnvStr()"}
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("*** FATAL: couldn't key value for variable '%s'", key)
		return "", e.BuildErr(errors.New("GetEnvStr() error"))
	}
	return val, nil
}

func GetEnvInt(key string) (int, error) {
	e := geterr.Err{Func: "GetEnvInt()"}
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("*** FATAL: couldn't key value for variable '%s'", key)
		return 0, e.BuildErr(errors.New("GetEnvStr() error"))
	}

	// convert key from string to int
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		e.Msg = fmt.Sprintf("*** FATAL: couldn't key value for variable '%s'", key)
		return 0, e.BuildErr(errors.New("error converting to int"))
	}
	return valAsInt, nil
}
