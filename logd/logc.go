package logd

import "fmt"

// basic log only to console
func logc(msg string) {
	fmt.Println("** |", msg)
}
