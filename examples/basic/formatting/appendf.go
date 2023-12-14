package formatting

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// AppendF -- Go 1.19 から追加された fmt.Appendf() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/fmt@go1.19#Appendf
//   - https://dev.to/emreodabas/quick-guide-go-119-features-1j40
func AppendF() error {
	var (
		buf = make([]byte, 0)
		sli = []any{"hello", "world"}
	)

	buf = fmt.Appendf(buf, "%s %s", sli[1], sli[0])
	output.Stdoutl("[fmt.Appendf]", string(buf))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: formatting_appendf

	   [Name] "formatting_appendf"
	   [fmt.Appendf]        world hello


	   [Elapsed] 20.19µs
	*/

}
