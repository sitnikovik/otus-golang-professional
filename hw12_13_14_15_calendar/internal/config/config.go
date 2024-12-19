package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// getEnv returns the environment variable by the key.
func getEnv(key string) string {
	return os.Getenv(key)
}

// Config describes the app configuration.
type Config struct {
	// Logger describes the logging configuration.
	Logger LoggerConf `yaml:"logger"`
	// Redis describes the redis configuration.
	Redis RedisConf `yaml:"redis"`
	// PG describes the postgres configuration.
	PG PGConf `yaml:"pg"`
	// HTTP describes the HTTP server configuration.
	HTTP HTTPConf `yaml:"http"`
}

// LoggerConf describes the logging configuration.
type LoggerConf struct {
	Level string `yaml:"level"`
}

// RedisConf describes the redis configuration.
type RedisConf struct {
	// Network describes the network type.
	Network string `yaml:"network"`
	// Addr describes the redis address.
	Addr string `yaml:"addr"`
	// Port describes the redis port.
	Port string `yaml:"port"`
}

// PGConf describes the postgres configuration.
type PGConf struct {
	// User describes the postgres user.
	User string `yaml:"user"`
	// Password describes the postgres password.
	Password string `yaml:"password"`
	// Database describes the postgres database.
	Database string `yaml:"database"`
	// Host describes the postgres host.
	Host string `yaml:"host"`
	// Port describes the postgres port.
	Port string `yaml:"port"`
}

// HTTPConf describes the HTTP server configuration.
type HTTPConf struct {
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}

// NewConfig creates and returns the new configuration.
func NewConfig(pathToFile string) (Config, error) {
	if err := godotenv.Load(pathToFile); err != nil {
		return Config{}, err
	}

	cfg := Config{}

	cfg.Logger.Level = getEnv("LOGGER_LEVEL")

	cfg.Redis.Network = getEnv("REDIS_NETWORK")
	cfg.Redis.Addr = getEnv("REDIS_ADDR")
	cfg.Redis.Port = getEnv("REDIS_PORT")

	cfg.PG.User = getEnv("PG_USER")
	cfg.PG.Password = getEnv("PG_PASSWORD")
	cfg.PG.Database = getEnv("PG_DATABASE")
	cfg.PG.Host = getEnv("PG_HOST")
	cfg.PG.Port = getEnv("PG_PORT")

	cfg.HTTP.Port = getEnv("HTTP_PORT")
	ttl, _ := strconv.Atoi(getEnv("HTTP_TIMEOUT"))
	cfg.HTTP.Timeout = ttl

	return cfg, nil
}
