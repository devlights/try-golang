package flags

import (
	"flag"
	"strings"

	"github.com/devlights/gomy/output"
)

// Func は、flag.Func() のサンプルです.
//
// flag.Func() は、func(string) error の関数を設定することで
// 任意の処理を行うことが出来る関数です。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#Func
func Func() error {
	var (
		hs []string
		fs = flag.NewFlagSet("", flag.ExitOnError)
	)

	fs.Func("hosts", "host names. comma separated.", func(v string) error {
		hs = strings.Split(v, ",")
		return nil
	})

	fs.Parse([]string{"-hosts", "example.invalid,my.local,localhost"})

	for _, h := range hs {
		output.Stdoutl("[h]", h)
	}

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: flags_func

		[Name] "flags_func"
		[h]                  example.invalid
		[h]                  my.local
		[h]                  localhost


		[Elapsed] 42.06µs
	*/

}
