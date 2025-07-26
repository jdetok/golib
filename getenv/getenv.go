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
	e := geterr.InitErr()
	if err := godotenv.Load(); err != nil {
		e.Msg = "failed to load .env variabels"
		return e.BuildErr(err)
	}
	return nil
}

func GetEnvStr(key string) (string, error) {
	e := geterr.InitErr()
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("key '%s' not found in .env", key)
		return "", e.BuildErr(errors.New(e.Msg))
	}
	return val, nil
}

func GetEnvInt(key string) (int, error) {
	e := geterr.InitErr()
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("key '%s' not found in .env", key)
		return 0, e.BuildErr(errors.New(e.Msg))
	}

	// convert key from string to int
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		e.Msg = fmt.Sprintf("couldn't convert '%s' value to int", key)
		return 0, e.BuildErr(errors.New(e.Msg))
	}
	return valAsInt, nil
}
