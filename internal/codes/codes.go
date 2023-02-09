// Package codes The commands exit codes.
package codes

// Code an int wrapper to represent the error codes.
type Code int

const (
	// UnknownError ...
	UnknownError Code = 1000

	// ReadError ...
	ReadError Code = 1001
	// UnmarshalError ...
	UnmarshalError Code = 1002

	// ConfigurationError ...
	ConfigurationError Code = 1003
)
