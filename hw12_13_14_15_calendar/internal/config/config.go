package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config describes the app configuration
type Config struct {
	// Logger describes the logging configuration
	Logger LoggerConf `yaml:"logger"`
	// Redis describes the redis configuration
	Redis RedisConf `yaml:"redis"`
	// PG describes the postgres configuration
	PG PGConf `yaml:"pg"`
	// HTTP describes the HTTP server configuration
	HTTP HTTPConf `yaml:"http"`
}

// LoggerConf describes the logging configuration
type LoggerConf struct {
	Level string `yaml:"level"`
}

// RedisConf describes the redis configuration
type RedisConf struct {
	// Network describes the network type
	Network string `yaml:"network"`
	// Addr describes the redis address
	Addr string `yaml:"addr"`
}

// PGConf describes the postgres configuration
type PGConf struct {
	// User describes the postgres user
	User string `yaml:"user"`
	// Password describes the postgres password
	Password string `yaml:"password"`
	// Database describes the postgres database
	Database string `yaml:"database"`
	// Host describes the postgres host
	Host string `yaml:"host"`
	// Port describes the postgres port
	Port string `yaml:"port"`
}

// HTTPConf describes the HTTP server configuration
type HTTPConf struct {
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
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
