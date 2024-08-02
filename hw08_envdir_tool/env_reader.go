package main

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"strings"
)

var (
	// ErrNilDirEntry says that dir entry is nil
	ErrNilDirEntry = errors.New("dirEntry is nil")
	// ErrInvalidEnvName says that env name contains invalid characters
	ErrInvalidEnvName = errors.New("invalid env name")
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)
	for _, dirEntry := range dirEntries {
		err := validateDirEntry(dirEntry)
		if err != nil {
			return nil, err
		}

		envName := dirEntry.Name()

		bb, err := os.ReadFile(dir + "/" + envName)
		if err != nil {
			return nil, err
		}

		env[envName] = EnvValue{
			Value:      prepareEnvValue(bb),
			NeedRemove: needRemoveEnv(envName),
		}
	}

	return env, nil
}

// validateDirEntry checks if dirEntry could be used as env variable
func validateDirEntry(dirEntry fs.DirEntry) error {
	if dirEntry == nil {
		return ErrNilDirEntry
	}
	if dirEntry.IsDir() {
		return nil
	}

	envName := dirEntry.Name()
	if strings.Contains(envName, "=") {
		return ErrInvalidEnvName
	}

	return nil
}

// prepareEnvValue prepares env value from file content
func prepareEnvValue(bb []byte) string {
	isToShiftString := !bytes.Contains(bb, []byte{0x00})
	bb = bytes.ReplaceAll(bb, []byte{0x00}, []byte{'\n'})

	s := string(bb)
	if isToShiftString {
		// shift string to the first '\n'
		for i := 0; i < len(s); i++ {
			if s[i] == '\n' {
				// there is no need to shift with strings.Split()
				s = s[:i]
				break
			}
		}
	}
	s = strings.TrimRight(s, " ")
	s = strings.Trim(s, "\t")

	return s
}

// needRemoveEnv checks if env variable is already set
func needRemoveEnv(envName string) bool {
	_, ok := os.LookupEnv(envName)
	return ok
}
