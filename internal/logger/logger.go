// Package logger The logger support for the CLI app.
package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Setup This function configures the logger.
func Setup(logLevel string) error {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
