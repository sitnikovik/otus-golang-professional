package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config describes the app configuration
type Config struct {
	// Logger describes the logging configuration
	Logger LoggerConf
}

// Load parses the app configuration with provided path and returns it
func Load(path string) (conf Config, err error) {
	bb, err := os.ReadFile(path)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(bb, &conf); err != nil {
		return
	}

	return
}
