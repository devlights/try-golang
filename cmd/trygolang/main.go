package main

import (
	"fmt"
	"log"
	"os"

	"github.com/devlights/try-golang/builder"
	"github.com/devlights/try-golang/mappings"
)

func main() {
	var (
		args    *Args
		mapping mappings.ExampleMapping
	)

	args = NewArgs()
	args.Parse()

	mapping = builder.BuildMappings()

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
