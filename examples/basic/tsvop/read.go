package tsvop

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/devlights/gomy/output"
)

// Read は、csv.Reader を利用したTSVデータの読み取りのサンプルです.
func Read() error {
	var (
		buf = new(bytes.Buffer)
		err error
	)

	fmt.Fprintln(buf, "hello\tworld")
	fmt.Fprintln(buf, "world\thello")

	var (
		r   = csv.NewReader(buf)
		rec []string
	)

	// TSV なので、区切り文字をタブに変更
	r.Comma = '\t'

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
}
