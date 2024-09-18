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
		if arg == "" {
			continue
		}

		// Parse flags
		flagParsed, err := parseFlag(ff, arg)
		if err != nil {
			return nil, nil, err
		}
		if flagParsed {
			continue
		}

		// Parse arguments
		err = parseArg(aa, arg)
		if err != nil {
			return nil, nil, err
		}
	}

	return aa, ff, nil
}

func parseFlag(flags *Flags, in string) (bool, error) {
	if !strings.Contains(in, "=") {
		return false, nil
	}

	parts := strings.SplitN(in, "=", 2)
	key, val := parts[0], parts[1]

	if key == "--timeout" {
		val = strings.TrimSuffix(val, "s")
		ttl, err := strconv.Atoi(val)
		if err != nil {
			return false, fmt.Errorf("parse timeout flag err: %w", err)
		}
		flags.Timeout = ttl
		return true, nil
	}

	return false, nil
}

func parseArg(args *Args, in string) error {
	if args.Address == "" {
		args.Address = in
		return nil
	}

	port, err := strconv.Atoi(in)
	if err != nil {
		return fmt.Errorf("port err: %w", err)
	}
	args.Port = port
	return nil
}
