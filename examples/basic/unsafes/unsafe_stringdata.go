package unsafes

import (
	"unsafe"

	"github.com/devlights/gomy/output"
)

// UnsafeStringData は、unsafe.StringData() のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/unsafe@go1.22.2#StringData
//   - https://mattn.kaoriya.net/software/lang/go/20220907112622.htm
func UnsafeStringData() error {
	const str = "hello world"

	buf := unsafe.StringData(str)
	str2 := unsafe.String(buf, len(str))

	output.Stdoutl("[str2]", str2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: unsafe_stringdata

	   [Name] "unsafe_stringdata"
	   [str2]               hello world


	   [Elapsed] 10.99µs
	*/

}
