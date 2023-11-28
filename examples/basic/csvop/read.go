package csvop

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/devlights/gomy/output"
)

// Read は、csv.Reader を利用したCSVデータの読み込みのサンプルです.
func Read() error {
	var (
		buf = new(bytes.Buffer)
		err error
	)

	fmt.Fprintln(buf, "hello,world")
	fmt.Fprintln(buf, "world,hello")

	var (
		r   = csv.NewReader(buf)
		rec []string
	)

LOOP:
	for {
		if rec, err = r.Read(); err != nil {
			switch err {
			case io.EOF:
				break LOOP
			default:
				return err
			}
		}

		output.Stdoutf("[rec]", "%v\tlen=%d\n", rec, len(rec))
	}

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: csv_read

	   [Name] "csv_read"
	   [rec]                [hello world]      len=2
	   [rec]                [world hello]      len=2


	   [Elapsed] 38.33µs
	*/

}
