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
	buf.Write([]byte("hello"))
	buf.WriteTo(os.Stdout)

	return nil
}
