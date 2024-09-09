package main

import (
	"errors"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}

	/*
	   $ task
	   task: [default] go run main.go
	   データが渡されていません。標準入力からデータを入力してください。
	   exit status 1
	   task: [default] echo helloworld | go run main.go
	   OK
	   task: [default] echo helloworld > test.txt
	   task: [default] go run main.go < test.txt
	   OK
	   task: [default] rm -f test.txt
	*/
}

func run() error {
	//
	// 標準入力からデータが渡されているかどうかのチェック
	//
	stat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return errors.New("データが渡されていません。標準入力からデータを入力してください。")
	}

	log.Println("OK")

	return nil
}
