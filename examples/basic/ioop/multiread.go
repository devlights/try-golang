package ioop

import (
	"fmt"
	"io"
	"os"
)

// MultiRead は、io.MultiReaderを利用して複数のファイルを一気に読み込むサンプルです。
//
// > MultiReader returns a Reader that's the logical concatenation of the provided input readers.
// They're read sequentially. Once all inputs have returned EOF, Read will return EOF.
// If any of the readers return a non-nil, non-EOF error, Read will return that error.
//
// > MultiReaderは、与えられた入力リーダーを論理的に連結したリーダーを返します。
// 順次読み込まれ、すべての入力がEOFを返したらEOFを返します。
// いずれかのリーダがEOF以外のエラーを返した場合、Readはそのエラーを返します。
//
// # REFERENCES
//
//   - https://pkg.go.dev/io@go1.22.2#MultiReader
func MultiRead() error {
	const (
		F1 = "go.mod"
		F2 = ".gitpod.yml"
		F3 = "main.go"
	)

	var (
		f1, f2, f3 *os.File
		err        error
	)

	if f1, err = os.Open(F1); err != nil {
		return err
	}
	defer f1.Close()

	if f2, err = os.Open(F2); err != nil {
		return err
	}
	defer f2.Close()

	if f3, err = os.Open(F3); err != nil {
		return err
	}
	defer f3.Close()

	var (
		r   = io.MultiReader(f1, f2, f3)
		buf []byte
	)

	buf, err = io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println(string(buf))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: ioop_multiread

	   [Name] "ioop_multireader"
	   module github.com/devlights/try-golang

	   go 1.22

	   require (
	           github.com/devlights/gomy v0.6.0
	           github.com/pelletier/go-toml/v2 v2.1.1
	           github.com/shopspring/decimal v1.3.1
	           golang.org/x/crypto v0.21.0
	           golang.org/x/exp v0.0.0-20240222234643-814bf88cf225
	           golang.org/x/sync v0.6.0
	           golang.org/x/term v0.18.0
	           golang.org/x/text v0.14.0
	           gopkg.in/ini.v1 v1.67.0
	           gopkg.in/natefinch/lumberjack.v2 v2.2.1
	           gopkg.in/yaml.v3 v3.0.1
	   )

	   require golang.org/x/sys v0.18.0 // indirect
	   image:
	     file: .gitpod.Dockerfile

	   tasks:
	     - name: initial
	       init:
	         go install github.com/go-task/task/v3/cmd/task@latest &&
	         go install honnef.co/go/tools/cmd/staticcheck@latest &&
	         go install golang.org/x/tools/cmd/goimports@latest &&
	         go install github.com/go-delve/delve/cmd/dlv@latest &&
	         task build
	       command:
	         go version


	   vscode:
	     extensions:
	       - golang.go
	       - TakumiI.markdowntable
	   package main

	   import "github.com/devlights/try-golang/cmd"

	   func main() {
	           cmd.Execute()
	   }



	   [Elapsed] 1.06411ms
	*/

}
