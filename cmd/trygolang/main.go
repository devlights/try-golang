package main

import (
	"fmt"
	"github.com/devlights/try-golang/interfaces"
	"github.com/devlights/try-golang/lib"
	"log"
	"os"
)

func main() {
	var (
		args    *Args
		mapping interfaces.SampleMapping
	)

	args = NewArgs()
	args.Parse()

	mapping = lib.MakeMapping()

	if args.ShowNames {
		printAllExampleNames(mapping)
		os.Exit(0)
	}

	defer fmt.Println("END")

	var cmd Command
	if args.ExampleName != "" {
		cmd = NewRunOnceCommand(NewRunOnceArgs(args.ExampleName, mapping))
	} else {
		cmd = NewRunLoopCommand(NewRunLoopArgs(args, mapping))
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
