package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

	if args.RunWithVsCode {
		if _, err := os.Stat(".target"); os.IsNotExist(err) {
			log.Println("VSCode 経由で実行する場合は .target ファイルが必要です")
			log.Println("(.target ファイルの中に実行したいサンプル名を入れてください)")
			log.Fatal("終了します...")
		}

		b, err := ioutil.ReadFile(".target")
		if err != nil {
			log.Fatalf("Cannot read .target file")
		}

		args.ExampleName = strings.TrimRight(string(b), "\n")
	}

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
