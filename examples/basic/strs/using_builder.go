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
}
