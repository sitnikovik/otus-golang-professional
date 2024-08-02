package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	returnCode = 0

	if len(cmd) == 0 {
		return
	}

	// Set environment variables
	for k, v := range env {
		if v.NeedRemove {
			os.Unsetenv(k)
		}
		os.Setenv(k, v.Value)
	}

	// Run command
	args := []string{}
	if len(cmd) > 1 {
		args = cmd[1:]
	}
	command := exec.Command(cmd[0], args...)
	bb, err := command.Output()
	if err != nil {
		returnCode = 1
	}
	fmt.Print(string(bb))

	return
}
