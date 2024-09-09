package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Args describes the arguments of the program.
type Args struct {
	Address string // The address of the server.
	Port    int    // The port of the server.
}

// Flags describes the flags of the program.
type Flags struct {
	Timeout int // The timeout of the connection.
}

// ParseInput parses the input arguments and flags.
func ParseInput(args []string) (*Args, *Flags, error) {
	aa, ff := &Args{}, &Flags{}

	for i := 1; i < len(args); i++ {
		arg := args[i]

		// Parse arguments
		if i < 3 && !strings.Contains(arg, "=") {
			if i == 1 {
				aa.Address = arg
			}
			if i == 2 {
				port, err := strconv.Atoi(arg)
				if err != nil {
					return nil, nil, fmt.Errorf("port err: %w", err)
				}
				aa.Port = port
			}

			continue
		}

		// Parse flags
		if strings.Contains(arg, "=") {
			parts := strings.SplitN(arg, "=", 2)
			key, value := parts[0], parts[1]

			if key == "--timeout" {
				ttl, err := strconv.Atoi(value)
				if err != nil {
					return nil, nil, fmt.Errorf("timeout err: %w", err)
				}
				ff.Timeout = ttl
			}
		}
	}

	return aa, ff, nil
}
