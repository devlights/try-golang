package error_

import (
	"fmt"
	"log"
	"os"
)

// error のハンドリングについてのサンプル
// REFERENCES::http://bit.ly/2IqmFD5, http://bit.ly/2IpThgk, http://bit.ly/2IpTzUs
func Error01() error {
	// 最初に error を宣言しておいて、ハンドリングする関数を定義し呼び出すようにするパターン
	var err error
	check := func() {
		if err != nil {
			log.Printf(err.Error())
		}
	}

	// GOPATH 取得
	gopath := os.Getenv("GOPATH")

	// ディレクトリ表示
	dir, err := os.Open(gopath)
	check()

	if dir == nil {
		log.Fatal("dir is nil")
	}

	names, err := dir.Readdirnames(0)
	check()

	for _, n := range names {
		fmt.Println(n)
	}

	return nil
}
