package unsafes

import (
	"syscall"
	"unsafe"

	"github.com/devlights/gomy/output"
)

// UnsafeString は、unsafe.String() のサンプルです.
//
// # REFERENCES
//
//   - https://pkg.go.dev/unsafe@go1.22.2#String
//   - https://mattn.kaoriya.net/software/lang/go/20220907112622.htm
func UnsafeString() error {
	const str = "hello world"

	//
	// *byte を作成
	//
	var (
		buf *byte
		err error
	)

	buf, err = syscall.BytePtrFromString(str)
	if err != nil {
		return err
	}

	//
	// unsafe.String() で、Goのstring型に変換
	//
	str2 := unsafe.String(buf, len(str))

	output.Stdoutl("[str ]", str)
	output.Stdoutl("[str2]", str2)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: unsafe_string

	   [Name] "unsafe_string"
	   [str ]               hello world
	   [str2]               hello world


	   [Elapsed] 12.471µs
	*/

}
