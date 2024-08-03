package main

import (
	"log"
	"os"
)

const (
	minNArgs = 2
)

var ErrorWrongNArgs = "wrong number of arguments, expected at least %d"

func main() {
	if len(os.Args)-1 < minNArgs {
		log.Fatalf(ErrorWrongNArgs, minNArgs)
	}

	env, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(RunCmd(os.Args[2:], env))
}
