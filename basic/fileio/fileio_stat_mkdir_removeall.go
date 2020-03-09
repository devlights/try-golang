package fileio

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// StatMkdirRemoveAll は、ディレクトリの存在確認と作成および削除のサンプルです.
func StatMkdirRemoveAll() error {
	// ディレクトリパスを生成
	dname := "try-golang-fileio03"
	dpath := filepath.Join(os.TempDir(), dname)

	fmt.Printf("ディレクトリ： %s\n", dpath)

	// 存在するか
	if _, err := os.Stat(dpath); os.IsNotExist(err) {
		fmt.Println("ディレクトリはまだ存在しない")
	}

	// ディレクトリ作成
	if err := os.MkdirAll(dpath, 0777); err != nil {
		log.Fatal(err)
		return err
	}

	// 存在するか
	if _, err := os.Stat(dpath); err == nil {
		fmt.Println("ディレクトリは存在する")
	}

	// 関数抜ける際にディレクトリ削除
	defer func() {
		if err := os.RemoveAll(dpath); err != nil {
			log.Fatal(err)
		}
	}()

	// ファイルを作成
	fname := "testfile"
	fpath := filepath.Join(dpath, fname)
	file, err := os.Create(fpath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// データ書き込み
	if _, err = file.WriteString("hello world"); err != nil {
		log.Fatal(err)
		return err
	}

	// 閉じる
	if err = file.Close(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
