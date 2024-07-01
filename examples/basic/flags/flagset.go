package flags

import (
	"flag"

	"github.com/devlights/gomy/output"
)

// Flagset は、flag.Flagsetのサンプルです。
//
// 通常、flagパッケージはアプリケーションの引数を扱うものであるが
// flag.Flagsetを利用することで、外から引数を指定してパースすることが可能となる。
// ユニットテストと相性が良い。
//
// # REFERENCES
//   - https://pkg.go.dev/flag@go1.22.4#FlagSet
func Flagset() error {
	var (
		fs       = flag.NewFlagSet("", flag.ExitOnError)
		intValue = fs.Int("intvalue", 0, "int value")
		strValue = fs.String("strvalue", "", "string value")
	)

	// Flagsetは自分でParse()を呼ぶ必要がある
	// アプリケーション引数を渡す場合は os.Args[1:] となる
	fs.Parse([]string{"-intvalue", "100", "-stvalue", "hello world"})

	output.Stdoutl("[intValue]", *intValue)
	output.Stdoutl("[intValue]", *strValue)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: flags_flagset

	   [Name] "flags_flagset"
	   [intValue]           100
	   [intValue]           hello world

	   [Elapsed] 43.41µs
	*/
}
