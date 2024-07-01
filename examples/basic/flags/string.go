package flags

import (
	"flag"

	"github.com/devlights/gomy/output"
)

// String は、flag.String(), flag.StringVar() のサンプルです。
//
// flagパッケージの関数は、flag.String()のように受け皿を戻り値で返してくれる関数と
// flag.StringVar() のように予め自前で用意している変数を利用する２パターンの使い方がある。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#String
//   - https://pkg.go.dev/flag@go1.22.4#StringVar
func String() error {
	var (
		fs = flag.NewFlagSet("", flag.ExitOnError)

		s1 *string
		s2 string
	)

	s1 = fs.String("s1", "", "string value 1")
	fs.StringVar(&s2, "i2", "default-value", "string value 2")

	fs.Parse([]string{"-s1", "helloworld"})

	output.Stdoutl("[s1]", *s1)
	output.Stdoutl("[s2]", s2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: flags_string

	   [Name] "flags_string"
	   [s1]                 helloworld
	   [s2]                 default-value


	   [Elapsed] 16.98µs
	*/

}
