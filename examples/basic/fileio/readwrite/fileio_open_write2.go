package readwrite

import (
	"fmt"
	"io"
	"os"

	"github.com/devlights/gomy/output"
)

// OpenWrite2 は、os.Create を使ったサンプルです.
func OpenWrite2() error {
	const (
		fpath = "/tmp/trygolang-openwrite2.txt"
	)

	var (
		f   io.WriteCloser
		err error
	)

	if f, err = os.Create(fpath); err != nil {
		return err
	}

	func() {
		defer f.Close()

		fmt.Fprintln(f, "\nhello world")
		fmt.Fprintln(f, "world hello")
	}()

	var (
		data []byte
	)

	if data, err = os.ReadFile(fpath); err != nil {
		return err
	}

	output.Stdoutl("[data]", string(data))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_open_write2

	   [Name] "fileio_open_write2"
	   [data]
	   hello world
	   world hello



	   [Elapsed] 88.26µs
	*/

}
