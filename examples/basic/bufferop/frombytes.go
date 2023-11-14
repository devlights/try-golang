package bufferop

import (
	"bytes"
	"os"
)

// FromBytes -- bytes.Buffer を バイト列 から生成するサンプルです.
func FromBytes() error {
	// bytes.NewBuffer() を利用すると
	// 指定したバイト列を初期値としたバッファを作成できる
	var (
		//                           H   E   L   L   O   \n
		buf = bytes.NewBuffer([]byte{72, 69, 76, 76, 79, 10})
	)

	if _, err := buf.WriteTo(os.Stdout); err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: bufferop_from_bytes

	   [Name] "bufferop_from_bytes"
	   HELLO


	   [Elapsed] 2.52µs
	*/

}
