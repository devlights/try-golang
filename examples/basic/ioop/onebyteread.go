package ioop

import (
	"bytes"
	"io"

	"github.com/devlights/gomy/output"
)

// OneByteRead は、１バイトずつ読み出す io.LimitedReader のサンプルです.
//
// 通信データのように固定部分を決まったサイズで読み取るときなどに便利。
//
// # REFERENCES
//   - https://pkg.go.dev/io@go1.19.3#LimitedReader
func OneByteRead() error {
	const (
		ReadSize = 0x01
		BufSize  = 0xff
	)

	var (
		message = "hello world"
		src     = bytes.NewBufferString(message)
	)

	for {
		var (
			reader = io.LimitReader(src, ReadSize)
			buf    = make([]byte, BufSize)
			size   int
			err    error
		)

		size, err = reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		output.Stdoutf("[OneByte]", "%d byte read: %v\n", size, string(buf[:size]))
	}

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: ioop_onebyte_read

	   [Name] "ioop_onebyte_read"
	   [OneByte]            1 byte read: h
	   [OneByte]            1 byte read: e
	   [OneByte]            1 byte read: l
	   [OneByte]            1 byte read: l
	   [OneByte]            1 byte read: o
	   [OneByte]            1 byte read:
	   [OneByte]            1 byte read: w
	   [OneByte]            1 byte read: o
	   [OneByte]            1 byte read: r
	   [OneByte]            1 byte read: l
	   [OneByte]            1 byte read: d


	   [Elapsed] 100.72µs
	*/

}
