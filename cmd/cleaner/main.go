package main

import (
	"os"
	"fmt"
	logger "project-cleaner-CLI/internal/util/logger"
)

func main() {
	var args []string = os.Args
	validateArgsLength(args);
	fmt.Println("Hello from cleaner!")
}

func validateArgsLength(args []string) {
	if len(args) < 2 {
		logger.Fatal("you must specify path(s) to clean")
	}
}