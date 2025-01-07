package config

import (
	"os"
	"path/filepath"
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
	// GRPC describes the GRPC server configuration.
	GRPC GRPCConf `yaml:"grpc"`
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
	// Host describes the HTTP server host.
	Host string `yaml:"host"`
	// Port describes the HTTP server port.
	Port string `yaml:"port"`
	// Timeout describes the HTTP server timeout.
	Timeout int `yaml:"timeout"`
}

// GRPCConf describes the GRPC server configuration.
type GRPCConf struct {
	// Host describes the GRPC server host.
	Host string `yaml:"host"`
	// Port describes the GRPC server port.
	Port string `yaml:"port"`
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

	cfg.GRPC.Port = getEnv("GRPC_PORT")
	cfg.GRPC.Host = getEnv("GRPC_HOST")

	return cfg, nil
}

// NewTestConfig creates and returns the new test configuration.
func NewTestConfig() (Config, error) {
	rootdir, err := findRootDir()
	if err != nil {
		return Config{}, err
	}
	envFile := filepath.Join(rootdir, ".env")

	return NewConfig(envFile)
}

// findRootDir searches for the root directory of the project by looking for go.mod file.
func findRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", os.ErrNotExist
		}
		dir = parentDir
	}
}
