package formatting

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// AppendLn -- Go 1.19 から追加された fmt.Appendln() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/fmt@go1.19#Appendf
//   - https://dev.to/emreodabas/quick-guide-go-119-features-1j40
func AppendLn() error {
	var (
		buf = make([]byte, 0)
		sli = []any{"hello", "world"}
	)

	for _, v := range sli {
		buf = fmt.Appendln(buf, v)
	}

	output.Stdoutl("[fmt.Appendln]", string(buf))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: formatting_appendln

	   [Name] "formatting_appendln"
	   [fmt.Appendln]       hello
	   world



	   [Elapsed] 19.13µs
	*/

}
