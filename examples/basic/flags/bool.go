package flags

import (
	"flag"

	"github.com/devlights/gomy/output"
)

// Bool は、flag.Bool(), flag.BoolVar() のサンプルです。
//
// flagパッケージの関数は、flag.Bool()のように受け皿を戻り値で返してくれる関数と
// flag.BoolVar() のように予め自前で用意している変数を利用する２パターンの使い方がある。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#Bool
//   - https://pkg.go.dev/flag@go1.22.4#BoolVar
func Bool() error {
	var (
		fs = flag.NewFlagSet("", flag.ExitOnError)

		b1 *bool
		b2 bool
	)

	b1 = fs.Bool("b1", false, "bool value 1")
	fs.BoolVar(&b2, "b2", true, "bool value 2")

	fs.Parse([]string{"-b1", "-b2=false"})

	output.Stdoutl("[b1]", *b1)
	output.Stdoutl("[b2]", b2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: flags_bool

	   [Name] "flags_bool"
	   [b1]                 true
	   [b2]                 false


	   [Elapsed] 36.81µs
	*/

}
