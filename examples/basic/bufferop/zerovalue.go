package bufferop

import (
	"bytes"
	"os"
)

// ZeroValue -- bytes.Buffer を ゼロ値 で利用した場合のサンプルです.
func ZeroValue() error {
	var (
		buf bytes.Buffer
	)

	// bytes.Buffer は、ゼロ値の場合は有効な空のバッファを表す
	if _, err := buf.Write([]byte("hello")); err != nil {
		return err
	}

	if _, err := buf.WriteTo(os.Stdout); err != nil {
		return err
	}

	return nil
}
