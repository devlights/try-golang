package bufferop

import (
	"bytes"
	"os"
)

// FromString -- bytes.Buffer を 文字列 から生成するサンプルです.
func FromString() error {
	// bytes.NewBufferString() を利用すると
	// 指定した文字列を初期値としたバッファを作成できる.

	const (
		str = "hello world"
	)

	var (
		buf = bytes.NewBufferString(str)
	)

	buf.WriteTo(os.Stdout)

	return nil
}
