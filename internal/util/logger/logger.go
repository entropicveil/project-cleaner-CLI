package logger

import (
	"fmt"
	"os"
)

const (
	colorRed   = "\x1b[31m"
	colorReset = "\x1b[0m"
)

func Fatal(msg string) {
	fmt.Fprintf(os.Stderr, "%sfatal:%s %s\n", colorRed, colorReset, msg)
	os.Exit(1)
}