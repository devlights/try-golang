package flags

import (
	"flag"

	"github.com/devlights/gomy/output"
)

// Int は、flag.Int(), flag.IntVar() のサンプルです。
//
// flagパッケージの関数は、flag.Int()のように受け皿を戻り値で返してくれる関数と
// flag.IntVar() のように予め自前で用意している変数を利用する２パターンの使い方がある。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#Int
//   - https://pkg.go.dev/flag@go1.22.4#IntVar
func Int() error {
	var (
		fs = flag.NewFlagSet("", flag.ExitOnError)

		i1 *int
		i2 int
	)

	i1 = fs.Int("i1", 0, "int value 1")
	fs.IntVar(&i2, "i2", 0, "int value 2")

	fs.Parse([]string{"-i1", "100", "-i2", "200"})

	output.Stdoutl("[i1]", *i1)
	output.Stdoutl("[i2]", i2)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: flags_int

	   [Name] "flags_int"
	   [i1]                 100
	   [i2]                 200


	   [Elapsed] 365.16µs
	*/

}
