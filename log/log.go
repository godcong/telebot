package log

import (
	"fmt"
)

var Debug = false

func Println(v ...interface{}) {
	if !Debug {
		return
	}
	fmt.Println(v...)
}

func Printfln(format string, v ...interface{}) {
	if !Debug {
		return
	}
	fmt.Sprintf(format, v)
	fmt.Println()
}
