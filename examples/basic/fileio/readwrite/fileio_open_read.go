package readwrite

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// OpenRead は、ファイルをOpenしてReadするサンプルです.
func OpenRead() error {
	// ファイルを操作する場合は os パッケージを利用する
	filename := "README.md"
	file, err := os.OpenFile(filename, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// GO言語ではIO周りのデータは文字列ではなくバイト列として扱うのみ
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return err
	}

	// バイト列を文字列にしたい場合は string([]byte) を使う
	fmt.Printf("%s", string(buf[:n]))

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: fileio_open_read

		[Name] "fileio_open_read"

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



		[Elapsed] 98.431µs
	*/

}
