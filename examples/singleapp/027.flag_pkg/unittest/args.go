package args

import (
	"flag"
	"io"
)

type (
	Value struct {
		A int
		B string
		C bool
	}
)

var (
	Empty = &Value{0, "", false}
	Error = &Value{1, "error", false}
)

func Parse(options []string) (*Value, error) {
	if len(options) == 0 {
		return Empty, nil
	}

	var (
		flags = flag.NewFlagSet("command", flag.ContinueOnError)
		a     = flags.Int("a", 0, "")
		b     = flags.String("b", "", "")
		c     = flags.Bool("c", false, "")
	)
	flags.SetOutput(io.Discard)

	err := flags.Parse(options)
	if err != nil {
		return Error, err
	}

	return &Value{*a, *b, *c}, nil
}
