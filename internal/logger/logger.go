package logger

import (
	"fmt"
	"os"
)

const prefix = "git-sage: "

func Fatal(format string, a ...any) int {
	fmt.Fprintf(os.Stderr, prefix+format+"\n", a...)
	return 1
}

func Info(format string, a ...any) {
	fmt.Fprintf(os.Stdout, prefix+format+"\n", a...)
}
