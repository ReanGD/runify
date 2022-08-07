package logger

import "fmt"

func Write(format string, a ...any) {
	fmt.Printf(format+"\n", a...)
}
