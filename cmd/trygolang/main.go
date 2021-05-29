package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/devlights/gomy/logops"
	"github.com/devlights/gomy/strops"
	"github.com/devlights/try-golang/builder"
	"github.com/devlights/try-golang/mapping"
	"github.com/devlights/try-golang/runner"
)

func main() {
	var (
		args    *Args
		mapping mapping.ExampleMapping
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

	var r runner.Runner
	if args.ExampleName != "" {
		r = runner.NewOnce(runner.NewOnceArgs(args.ExampleName, mapping))
	} else {
		r = runner.NewLoop(runner.NewLoopArgs(args.OneTime, mapping))
	}

	if err := r.Run(); err != nil {
		errLog.Fatal(err)
	}

	os.Exit(0)
}
