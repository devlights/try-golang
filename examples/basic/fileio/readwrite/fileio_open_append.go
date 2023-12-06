package readwrite

import (
	"fmt"
	"io"
	"os"

	"github.com/devlights/gomy/output"
)

// OpenAppend は、追記モードでファイルを開くサンプルです。
func OpenAppend() error {
	const (
		fpath = "/tmp/trygolang-openappend.txt"
	)

	write := func(fpath string) error {
		var (
			f     io.WriteCloser
			err   error
			flgs  int         = os.O_APPEND | os.O_CREATE | os.O_WRONLY
			fmode os.FileMode = 0644
		)

		if f, err = os.OpenFile(fpath, flgs, fmode); err != nil {
			return err
		}
		defer f.Close()

		fmt.Fprintln(f, "hello world")

		return nil
	}

	var (
		err error
	)

	// 2回書き込み
	for range make([]struct{}, 2) {
		if err = write(fpath); err != nil {
			return err
		}
	}

	// 結果確認
	var (
		r io.ReadCloser
	)

	if r, err = os.Open(fpath); err != nil {
		return err
	}
	defer r.Close()

	var (
		data []byte
	)

	if data, err = io.ReadAll(r); err != nil {
		return err
	}

	output.Stdoutl("[append]", string(data))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_open_append

	   [Name] "fileio_open_append"
	   [append]             hello world
	   hello world



	   [Elapsed] 135.58µs
	*/

}
