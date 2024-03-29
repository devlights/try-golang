package ioop

import (
	"bytes"
	"io"

	"github.com/devlights/gomy/output"
)

// LimitRead は、io.LimitedReader のサンプルです.
//
// 通信データのように固定部分を決まったサイズで読み取るときなどに便利。
//
// # REFERENCES
//   - https://pkg.go.dev/io@go1.19.3#LimitedReader
func LimitRead() error {
	const (
		ReadSize = 0x04
		BufSize  = 0xff
	)

	var (
		message     = "hello world"
		src         = bytes.NewBufferString(message)
		limitReader = io.LimitReader(src, ReadSize)
	)
	output.Stdoutf("[LimitReader]", "original: %v\n", message)

	for {
		var (
			buf  = make([]byte, BufSize)
			size int
			err  error
		)

		size, err = limitReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		output.Stdoutf("[LimitRead]", "%d byte(s) read: %v\n", size, string(buf[:size]))
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: ioop_limit_read

	   [Name] "ioop_limit_read"
	   [LimitReader]        original: hello world
	   [LimitRead]          4 byte(s) read: hell


	   [Elapsed] 22.76µs
	*/

}
