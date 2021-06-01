package fileio

import (
	"bytes"
	"io"
	"os"

	"github.com/devlights/gomy/output"
)

// OpenRead2 は、os.Openを使ったファイルを読み込みのサンプルです.
func OpenRead2() error {
	var (
		f   io.ReadCloser
		err error
	)

	if f, err = os.Open("README.md"); err != nil {
		return err
	}
	defer f.Close()

	var (
		data []byte
	)

	if data, err = io.ReadAll(f); err != nil {
		return err
	}

	line := data[:bytes.Index(data, []byte("\n"))]
	output.Stdoutl(string(line))

	return nil
}
