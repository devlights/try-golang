package strs

import (
	"strings"

	"github.com/devlights/gomy/output"
)

// UsingBuilder -- strings.Builder を利用したサンプルです.
func UsingBuilder() error {
	var (
		builder strings.Builder
	)

	builder.WriteString("hello")
	builder.WriteRune(' ')
	builder.WriteString("world")

	output.Stdoutl("builder", builder.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: string_using_builder

	   [Name] "string_using_builder"
	   builder              hello world


	   [Elapsed] 12.27µs
	*/

}
