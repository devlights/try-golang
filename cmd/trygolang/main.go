package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/devlights/gomy/logops"
	"github.com/devlights/gomy/strops"
	"github.com/devlights/try-golang/internal/builder"
	"github.com/devlights/try-golang/pkg/command"
	"github.com/devlights/try-golang/pkg/mappings"
)

func main() {
	var (
		args    *Args
		mapping mappings.ExampleMapping
	)

	appLog, errLog, _ := logops.Default.Logger(true, func(_, e, _ *log.Logger) {
		e.SetPrefix("[Error] ")
	})

	args = NewArgs()
	args.Parse()

	if args.RunWithVsCode {
		if _, err := os.Stat(".target"); os.IsNotExist(err) {
			appLog.Println("--------------------------------------------------------")
			appLog.Println("VSCode 経由で実行する場合は .target ファイルが必要です")
			appLog.Println("(.target ファイルの中に実行したいサンプル名を入れてください)")
			appLog.Println("例: $ echo 'helloworld' > .target ")
			appLog.Println("--------------------------------------------------------")
			appLog.Fatal("終了します...")
		}

		b, err := ioutil.ReadFile(".target")
		if err != nil {
			errLog.Fatalf("Cannot read .target file")
		}

		args.ExampleName = strops.Chop(string(b))
	}

	mapping = builder.BuildMappings()

	if args.ShowNames {
		appLog.Println("[Examples]")
		for _, v := range mapping.AllExampleNames() {
			appLog.Printf("\t%s", v)
		}

		os.Exit(0)
	}

	defer fmt.Println("END")

	var cmd command.Cmd
	if args.ExampleName != "" {
		cmd = NewRunOnceCommand(NewRunOnceArgs(args.ExampleName, mapping))
	} else {
		cmd = NewRunLoopCommand(NewRunLoopArgs(args, mapping))
	}

	if err := cmd.Run(); err != nil {
		errLog.Fatal(err)
	}

	os.Exit(0)
}
