package bufferop

import (
	"bytes"
	"encoding/csv"

	"github.com/devlights/gomy/output"
)

// UseAsReader -- bytes.Buffer を io.Writer として利用するサンプルです.
func UseAsWriter() error {
	// bytes.Buffer は io.Writer を実装しているので
	// io.Readerが必要な様々な場面で利用できる
	//
	// 注意点として、io.Writerを実装しているのは
	//   *bytes.Buffer
	// の方であるという点。
	//
	var (
		rec    = []string{"1", "2", "3"}
		buf    = new(bytes.Buffer)
		writer = csv.NewWriter(buf)
	)

	if err := writer.Write(rec); err != nil {
		return err
	}
	writer.Flush()

	output.Stdoutl("[buf]", buf.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: bufferop_use_as_writer

	   [Name] "bufferop_use_as_writer"
	   [buf]                1,2,3



	   [Elapsed] 22.4µs
	*/

}
