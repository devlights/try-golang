package effectivego

import (
	"fmt"
	"strings"
)

// Effective Go - Slices の 内容についてのサンプルです。
func Slices() error {
	/*
		https://golang.org/doc/effective_go.html#slices

		- スライスは内部にデータとなる配列の参照を保持している
		- スライスを関数の引数として渡す場合は少し注意が必要。
		  - basic/slices/slice_pointer.go を参照
		- スライスをパラメータに取る標準ライブラリは多い。(os.Fileなど)
		- 配列を ary[x:y] のようにスライシングすることでもスライスは取得できる
		- スライスに要素を追加するには append() を利用する
		- スライスの内容を別のスライスにコピーするには copy() を利用する
	*/
	// 最初の5バイトを1バイトずつ読み込み
	const (
		s = "helloworld"
	)

	var (
		buf       = make([]byte, len(s))
		readBytes = 0
		reader    = strings.NewReader(s)
	)

	for i := 0; i < 5; i++ {
		// 1バイト分のスライスをReadに指定することで1バイトずつ読み込んでいる
		n, err := reader.Read(buf[i : i+1])
		readBytes += n

		if n == 0 || err != nil {
			fmt.Println(err)
			break
		}
	}

	fmt.Printf("readBytes:%d\n", readBytes)
	fmt.Printf("[buf] str:%v\tbuf:%v\n", string(buf), buf)

	// スライスをコピーするには copy() を使う
	buf2 := make([]byte, readBytes)
	copy(buf2, buf[:readBytes])

	// 配列にコピーしたい場合は以下のようにする
	buf3 := [5]byte{}
	copy(buf3[:readBytes], buf[:readBytes])

	fmt.Printf("buf:%v\tbuf2:%v\tbuf3:%v\n", buf, buf2, buf3)

	return nil
}
