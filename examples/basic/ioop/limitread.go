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
		READ_SIZE = 0x04
		BUF_SIZE  = 0xff
	)

	var (
		message     = "hello world"
		src         = bytes.NewBufferString(message)
		limitReader = io.LimitReader(src, READ_SIZE)
	)
	output.Stdoutf("[LimitReader]", "original: %v\n", message)

	for {
		var (
			buf  = make([]byte, BUF_SIZE)
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
}
