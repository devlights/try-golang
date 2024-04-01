package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// Pid は、os.Getpid()のサンプルです.
//
// os.Getpid() は、呼び出し元のプロセスIDを返す。
// os.Getpid() は、errorを返さない。
//
// # REFENRECES
//   - https://pkg.go.dev/os@go1.22.1#Getpid
func Pid() error {
	pid := os.Getpid()
	output.Stdoutl("[pid]", pid)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_pid

	   [Name] "osop_pid"
	   [pid]                35548


	   [Elapsed] 23.09µs
	*/

}
