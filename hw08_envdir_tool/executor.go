package main

import (
	"errors"
	"os"
	"os/exec"
)

const (
	// OkCode is a return code for success
	OkCode = 0
	// ErrCode is a return code for errors
	ErrCode = 1
)

var (
	// ErrNilEnvironment says that environment is nil
	ErrNilEnvironment = errors.New("environment is nil")
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	var err error
	returnCode = OkCode

	defer func() {
		if err != nil {
			logError(err)
		}
	}()

	if env == nil {
		err = ErrNilEnvironment
		returnCode = ErrCode
		return
	}
	if err = setEnvironment(env); err != nil {
		returnCode = ErrCode
		return
	}
	if len(cmd) == 0 {
		return
	}

	args := []string{}
	if len(cmd) > 1 {
		args = cmd[1:]
	}
	if err = run(cmd[0], args); err != nil {
		returnCode = ErrCode
	}

	return
}

func logError(err error) {
	if err == nil {
		return
	}

	os.Stderr.WriteString(err.Error())
}

func setEnvironment(env Environment) (err error) {
	for k, v := range env {
		if v.NeedRemove && os.Unsetenv(k) != nil {
			return
		}
		if err = os.Setenv(k, v.Value); err != nil {
			return err
		}
	}

	return
}

func run(cmd string, args []string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}
