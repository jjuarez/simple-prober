// Package utils Misc utilities to use across commands.
package utils

import (
	"log"
	"os"

	"github.com/jjuarez/simple-prober/internal/codes"
)

// ExitCommand This function implements a simple way to exit from the commands.
func ExitCommand(exitCode codes.Code, err error) {
	if err != nil {
		log.Println(err)
	}
	os.Exit(int(exitCode))
}
