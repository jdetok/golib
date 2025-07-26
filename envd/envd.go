package envd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/jdetok/golib/errd"
	"github.com/joho/godotenv"
)

func LoadDotEnv() error {
	e := errd.InitErr()
	if err := godotenv.Load(); err != nil {
		e.Msg = "failed to load .env variabels"
		return e.BuildErr(err)
	}
	return nil
}

func EnvStr(key string) string {
	e := errd.InitErr()
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("key '%s' not found in .env", key)
		fmt.Println(e.NewErr())
	}
	return val
}

func EnvInt(key string) int {
	e := errd.InitErr()
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("key '%s' not found in .env", key)
		fmt.Println(e.NewErr())
	}

	// convert key from string to int
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		e.Msg = fmt.Sprintf("couldn't convert '%s' value to int", key)
		fmt.Println(e.NewErr())
	}
	return valAsInt
}

func GetEnvStr(key string) (string, error) {
	e := errd.InitErr()
	val, ok := os.LookupEnv(key)
	if !ok {
		e.Msg = fmt.Sprintf("key '%s' not found in .env", key)
		return "", e.BuildErr(errors.New(e.Msg))
	}
	return val, nil
}

func GetEnvInt(key string) (int, error) {
	e := errd.InitErr()
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
