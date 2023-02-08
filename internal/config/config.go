// Package config This package groups the functions related to the configuration.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jjuarez/simple-prober/internal/codes"
	"github.com/jjuarez/simple-prober/internal/model"
	yaml "gopkg.in/yaml.v3"
)

// LoadFile This function load the content of the configuration file.
func loadFile(configFileName string) ([]byte, error) {
	var err error
	var fileContent []byte
	var eConfigFileName string

	eConfigFileName, _ = filepath.Abs(configFileName)
	if _, err = os.Stat(eConfigFileName); err != nil {
		return nil, fmt.Errorf("the configuration file: %s does not exist(errno: %d)", eConfigFileName, codes.ReadError)
	}

	if fileContent, err = os.ReadFile(eConfigFileName); err != nil {
		return nil, fmt.Errorf("something went wrong reading from the configuration file: %s (errno: %d)", eConfigFileName, codes.ReadError)
	}
	return fileContent, nil
}

// Load This function is responsible for the YAML configuration loading process.
func Load(configFileName string) (*[]model.Endpoint, error) {
	var err error
	var endpoints []model.Endpoint
	var configContent []byte

	if configContent, err = loadFile(configFileName); err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(configContent, &endpoints); err != nil {
		return nil, err
	}
	return &endpoints, nil
}
