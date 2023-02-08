// Package codes The commands exit codes.
package codes

// Code an int wrapper to represent the error codes.
type Code int

const (
	// ReadError ...
	ReadError Code = 1001
	// UnmarshalError ...
	UnmarshalError Code = 1002
)
