package readwrite

import (
	"io"
	"os"

	"github.com/devlights/gomy/output"
)

// OpenRead2 は、os.Openを使ったファイルを読み込みのサンプルです.
func OpenRead2() error {
	var (
		f   io.ReadCloser
		err error
	)

	// os.Open() は、読込み専用でファイルを開く
	if f, err = os.Open("README.md"); err != nil {
		return err
	}
	defer f.Close()

	var (
		data []byte
	)

	if data, err = io.ReadAll(f); err != nil {
		return err
	}

	output.Stdoutl(string(data[:1024]))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_open_read2

	   [Name] "fileio_open_read2"

	   # try-golang

	   This is my TUTORIAL project for golang

	   ![try-golang - Go Version](https://img.shields.io/badge/go-1.21-blue.svg)
	   ![Go](https://github.com/devlights/try-golang/workflows/Go/badge.svg?branch=master)

	   [![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/devlights/try-golang)

	   ## Go version

	   ```shell script
	   $ lsb_release -a
	   No LSB modules are available.
	   Distributor ID: Ubuntu
	   Description:    Ubuntu 20.04.5 LTS
	   Release:        20.04
	   Codename:       focal

	   $ go version
	   go version go1.21.0 linux/amd64

	   $ task build
	   task: [build] go build .

	   $ go version try-golang
	   try-golang: go1.21.0
	   ```

	   ## Run

	   ```shell script
	   $ go run main.go
	   ```

	   If you want to use [go-task](https://github.com/go-task/task), type the following command.

	   ```sh
	   $ go install github.com/go-task/task/v3/cmd/task@latest
	   ```

	   Once the above command is complete, you can run it at

	   ```sh
	   $ task run
	   ```

	   ## Test

	   ```shell script
	   $ go test -v ./...
	   ```

	   or

	   ```shell script
	   $ task test
	   ```




	   [Elapsed] 101.6µs
	*/

}
