package main

import (
	"fmt"

	"github.com/jdetok/golib/getenv"
)

func main() {
	_, err := getenv.GetEnvStr("TEST")
	if err != nil {
		fmt.Println(err)
	}
}
