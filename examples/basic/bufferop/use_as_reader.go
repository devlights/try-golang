package bufferop

import (
	"bytes"
	"encoding/csv"

	"github.com/devlights/gomy/output"
)

// UseAsReader -- bytes.Buffer を io.Reader として利用するサンプルです.
func UseAsReader() error {
	// bytes.Buffer は io.Reader を実装しているので
	// io.Readerが必要な様々な場面で利用できる
	//
	// 注意点として、io.Readerを実装しているのは
	//   *bytes.Buffer
	// の方であるという点。
	//
	var (
		buf    = bytes.NewBufferString("1,2,3")
		reader = csv.NewReader(buf)
	)

	var (
		rec []string
		err error
	)

	if rec, err = reader.Read(); err != nil {
		return err
	}

	output.Stdoutl("[rec]", rec)

	return nil
}
