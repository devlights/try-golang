package tsvop

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/devlights/gomy/output"
)

// Write は、csv.Writer を利用したTSVデータの書き込みのサンプルです.
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

	// TSV なので、区切り文字を変更
	w.Comma = '\t'

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
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: tsv_write

	   [Name] "tsvop_write"
	   [data]
	   hello   world
	   world   hello


	   [Elapsed] 12.76µs
	*/

}
