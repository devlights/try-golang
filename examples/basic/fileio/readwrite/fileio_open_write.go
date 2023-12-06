package readwrite

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
)

// OpenWrite は、ファイルをOpenしてWriteするサンプルです.
func OpenWrite() error {
	// 一時ファイルの作成
	file, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("TmpFile: %s\n", file.Name())

	// 関数を抜けた際に、利用したファイルを削除
	defer func() {
		_ = os.Remove(file.Name())

		// ファイルの存在チェックは、Go言語ではこうやる
		// REF: http://bit.ly/2I1LzYa
		if _, err = os.Stat(file.Name()); errors.Is(err, fs.ErrNotExist) {
			fmt.Println("存在しない")
		}
	}()

	// ファイル存在するか確認
	_, err = os.Stat(file.Name())
	if err == nil {
		fmt.Println("ファイル存在する")
	}

	// データを書き込む
	message := "hello world"
	_, err = file.WriteString(message)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 閉じる
	err = file.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 読み出してみる
	data, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
		return err
	}

	message = string(data)
	fmt.Println(message)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_open_write

	   [Name] "fileio_open_write"
	   TmpFile: /tmp/example717995680
	   ファイル存在する
	   hello world
	   存在しない


	   [Elapsed] 214.18µs
	*/

}
