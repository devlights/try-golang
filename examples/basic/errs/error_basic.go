package errs

import (
	"fmt"
	"log"
	"os"
)

// Basic -- error のハンドリングについてのサンプル
// REFERENCES::
//   - https://github.com/robpike/ivy/blob/master/ivy_test.go
//   - https://stackoverflow.com/questions/9371031/how-do-i-create-crossplatform-file-paths-in-go
//   - https://qiita.com/andromeda/items/c5195307cd08537d4fad
func Basic() error {
	// 最初に error を宣言しておいて、ハンドリングする関数を定義し呼び出すようにするパターン
	var err error
	check := func() {
		if err != nil {
			log.Print(err.Error())
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

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: error_basic

	   [Name] "error_basic"
	   bin
	   pkg


	   [Elapsed] 73.89µs
	*/

}
