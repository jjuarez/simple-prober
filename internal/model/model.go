// Package model These are only the types used to marshall the config file.
package model

import (
	"context"
	"fmt"
	"net"
	"time"
)

type Stringable interface {
	String() string
}

type Printable interface {
	Println()
}

// Endpoint The endpoint basic data.
type Endpoint struct {
	Name    string `yaml:"name,omitempty"`
	Kind    string `yaml:"kind,omitempty"`
	Address string `yaml:"address"`
}

// New Creates an Endpoint.
func New(name string, kind string, address string) *Endpoint {
	return &Endpoint{
		Name:    name,
		Kind:    kind,
		Address: address,
	}
}

// String Implements the Stingable interface for Endpoint.
func (e Endpoint) String() string {
	return fmt.Sprintf("%T(name:%s, kind:%s, address:%s)", e, e.Name, e.Kind, e.Address)
}

// Println Implements the Printable interface for Endpoint.
func (e Endpoint) Println() {
	fmt.Println(e.String())
}

// TestConnection ...
func (e Endpoint) TestConnection(timeout time.Duration) (bool, error) {
	var dialer net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	connection, err := dialer.DialContext(ctx, "tcp", e.Address)
	if err != nil {
		return false, err
	}
	defer connection.Close()

	if _, err := connection.Write([]byte("Testing connection...")); err != nil {
		return false, err
	}
	return true, nil
}
