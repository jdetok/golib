package misc

import (
	"fmt"
	"reflect"
	"time"
)

// print names of fields in structs (receursive) - not used in mlb etl code
func printStructFields(t reflect.Type) {
	for i := range t.NumField() {
		f := t.Field(i)
		fmt.Println(f.Name)
		if f.Type.Kind() == reflect.Struct && f.Type != reflect.TypeOf(time.Time{}) {
			printStructFields(f.Type)
		}
	}
}
