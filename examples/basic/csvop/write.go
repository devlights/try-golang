package csvop

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/devlights/gomy/output"
)

// Write は、csv.Writer を利用したCSVデータの書き込みのサンプルです.
func Write() error {
	var (
		buf = new(bytes.Buffer)
		err error
	)

	var (
		w       = csv.NewWriter(buf)
		records = [][]string{
			{"hello", "world"},
			{"world", "hello"},
		}
	)

	for _, rec := range records {
		if err = w.Write(rec); err != nil {
			return err
		}
	}
	w.Flush()

	var (
		data []byte
	)

	if data, err = io.ReadAll(buf); err != nil {
		return err
	}

	output.Stdoutf("[data]", "\n%v", string(data))

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: csv_write

	   [Name] "csv_write"
	   [data]
	   hello,world
	   world,hello


	   [Elapsed] 120.53µs
	*/

}
