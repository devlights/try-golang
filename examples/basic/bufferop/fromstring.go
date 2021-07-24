package bufferop

import (
	"bytes"
	"os"
)

// FromString -- bytes.Buffer を 文字列 から生成するサンプルです.
func FromString() error {
	const (
		str = "hello world"
	)

	// bytes.NewBufferString() を利用すると
	// 指定した文字列を初期値としたバッファを作成できる.

	var (
		buf = bytes.NewBufferString(str)
	)

	if _, err := buf.WriteTo(os.Stdout); err != nil {
		return err
	}

	return nil
}
