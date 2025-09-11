/*
go test の -short オプションを付与した際のサンプルです。
詳細は lib/lib_test.go を参照ください。

REFERENCES:
  - https://dev.to/jonasbn/til-skipping-tests-in-go-3i5l
  - https://golang.org/cmd/go/#hdr-Testing_flags
  - https://golang.org/pkg/testing/#hdr-Skipping
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		exitCode int
		err      error
	)

	if err = run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		exitCode = 1
	}

	os.Exit(exitCode)
}

func run() error {
	return nil
}
