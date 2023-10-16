package util

import (
	"fmt"
	"os"
)

const (
	exitOK     int = 0
	exitError  int = 1
	exitCancel int = 2
)

// Returns true if err != nil
func CheckError(err error) bool {
	return err != nil
}

func CheckErrorFatal(err error, msg string) {
	if CheckError(err) {
		fmt.Println("fatal: " + msg)
		os.Exit(exitError)
	}
}
