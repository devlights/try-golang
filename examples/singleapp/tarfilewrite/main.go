/*
	archive/tar の サンプルです。

	REFERENCES:
	  - https://pkg.go.dev/archive/tar@latest
*/
package main

import (
	"archive/tar"
	"os"
	"time"
)

func _err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// tar.Writer を 取得
	tw := tar.NewWriter(os.Stdout)

	// tar.Header を書き込み
	fh := tar.Header{
		Name:    "test.txt",
		Size:    11,
		Mode:    0644,
		ModTime: time.Now().Truncate(24 * time.Hour),
	}

	err := tw.WriteHeader(&fh)
	_err(err)

	// データを書き込み
	_, err = tw.Write([]byte("hello world"))
	_err(err)

	// 最後に Close() を呼んでおかないとtarファイルが生成できないので注意
	_err(tw.Flush())
	_err(tw.Close())
}
