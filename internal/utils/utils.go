// Package utils Misc utilities to use across commands.
package utils

import (
	"os"

	"github.com/jjuarez/simple-prober/internal/codes"
	log "github.com/sirupsen/logrus"
)

// ExitCommand This function implements a simple way to exit from the commands.
func ExitCommand(exitCode codes.Code, err error) {
	if err != nil {
		log.Error(err)
	}
	os.Exit(int(exitCode))
}
