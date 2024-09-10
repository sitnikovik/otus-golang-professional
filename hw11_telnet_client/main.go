package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	args, flags, err := ParseInput(os.Args)
	if err != nil {
		log.Fatalf("parse input err: %v", err)
	}

	var addr string
	if args.Port != 0 {
		addr = fmt.Sprintf("%s:%d", args.Address, args.Port)
	} else {
		addr = args.Address
	}
	telnetClient := NewTelnetClient(
		addr,
		time.Duration(flags.Timeout)*time.Second,
		os.Stdin,
		os.Stdout,
	)

	telnetClient.Connect()
	defer telnetClient.Close()

	for {
		telnetClient.Send()
		telnetClient.Receive()
	}
}
