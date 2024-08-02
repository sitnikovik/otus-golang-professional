package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
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
	scandir, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)
	for _, file := range scandir {
		if file.IsDir() {
			continue
		}

		envName := file.Name()
		if strings.Contains(envName, "=") {
			return nil, fmt.Errorf("env name '%s' contains '='", envName)
		}

		bb, err := os.ReadFile(dir + "/" + envName)
		if err != nil {
			return nil, err
		}

		// терминальные нули (0x00) заменяются на перевод строки (\n)
		isToShiftString := !bytes.Contains(bb, []byte{0x00})
		bb = bytes.ReplaceAll(bb, []byte{0x00}, []byte{'\n'})

		s := string(bb)
		if isToShiftString {
			for i := 0; i < len(s); i++ {
				if s[i] == '\n' {
					s = s[:i]
					break
				}
			}
		}
		s = strings.Trim(strings.TrimRight(s, " "), "\t")

		_, ok := os.LookupEnv(envName)
		env[envName] = EnvValue{
			Value:      s,
			NeedRemove: ok,
		}
	}

	return env, nil
}
