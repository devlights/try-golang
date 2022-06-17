/*
	archive/zip の サンプルです。

	REFERENCES:
	  - https://pkg.go.dev/archive/zip@latest
*/
package main

import (
	"archive/zip"
	"bufio"
	"os"
)

func _err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// zip.Writer を 取得
	// 最後に Close() を呼んでおかないとZipファイルが生成できないので注意
	zw := zip.NewWriter(os.Stdout)

	// zipファイル内にファイルを追加
	w, err := zw.Create("test.txt")
	_err(err)

	bw := bufio.NewWriter(w)
	bw.WriteString("hello world")
	_err(bw.Flush())

	// zipファイル内にディレクトリ付きでファイルを追加
	w, err = zw.Create("dir1/test2.txt")
	_err(err)

	bw = bufio.NewWriter(w)
	bw.WriteString("world hello")
	_err(bw.Flush())

	// zipファイルコメントを設定
	err = zw.SetComment("this is test zip file")
	_err(err)

	_err(zw.Flush())
	_err(zw.Close())
}
