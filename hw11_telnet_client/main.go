package main

import (
	"log"
	"os"
)

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	args, flags, err := ParseInput(os.Args)
	if err != nil {
		log.Fatalf("parse input err: %v", err)
	}

	log.Println(args, flags)
}
