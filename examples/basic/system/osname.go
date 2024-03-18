package system

import (
	"runtime"

	"github.com/devlights/gomy/output"
)

// OsName は、OSの名前を出力するサンプルです.
func OsName() error {
	output.Stdoutl("[OS]", runtime.GOOS)
	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: os_name

	   [Name] "os_name"
	   [OS]                 linux


	   [Elapsed] 6.21µs
	*/

}
