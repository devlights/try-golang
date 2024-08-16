package strs

import (
	"strings"

	"github.com/devlights/gomy/output"
)

// SplitFields は、strings.Fields() のサンプルです.
//
// 空白で分割したい場合は、strings.Split() を利用するより
// strings.Fields() を利用したほうが楽。
//
// # REFERENCES
//   - https://pkg.go.dev/strings@go1.23.0#Fields
func SplitFields() error {
	var (
		s = "hello world こんにちは     世界"
		p = strings.Fields(s)
	)

	for _, v := range p {
		output.Stdoutl("[value]", v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: string_split_fields

	   [Name] "string_split_fields"
	   [value]              hello
	   [value]              world
	   [value]              こんにちは
	   [value]              世界


	   [Elapsed] 33.651µs
	*/

}
