package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// Setenv は、os.Setenv() のサンプルです。
//
// 既に存在する環境変数に対して os.Setenv() した場合は
// そのプロセス内で値が上書きされる。
//
// # REFERENCES
//
//   - https://pkg.go.dev/os@go1.22.0#Setenv
func Setenv() error {
	var err error

	err = os.Setenv("MYENV1", "HELLOWORLD")
	if err != nil {
		return err
	}

	output.Stdoutl("[MYENV1]", os.Getenv("MYENV1"))

	err = os.Setenv("HOSTNAME", "HELLOWORLD")
	if err != nil {
		return err
	}

	output.Stdoutl("[HOSTNAME]", os.Getenv("HOSTNAME"))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: osop_setenv

	   [Name] "osop_setenv"
	   [MYENV1]             HELLOWORLD
	   [HOSTNAME]           HELLOWORLD


	   [Elapsed] 61.65µs
	*/

}
