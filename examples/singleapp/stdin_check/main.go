package main

import (
	"errors"
	"io/fs"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// 標準入力からデータが渡されているかどうかのチェック
	//
	stat, err := os.Stdin.Stat()
	if err != nil {
		return err
	}

	mode := stat.Mode()
	show(mode)

	if (mode & os.ModeCharDevice) != 0 {
		return errors.New("データが渡されていません。標準入力からデータを入力してください。")
	}

	return nil
}

func show(mode fs.FileMode) {
	log.Printf("Mode:         %v", mode)
	log.Printf("Is Directory: %v", mode.IsDir())
	log.Printf("Is Regular:   %v", mode.IsRegular())
	log.Printf("Permissions:  %v", mode.Perm())
	log.Printf("Mode string:  %s", mode.String())
	log.Printf("Char device?  %v", (mode & os.ModeCharDevice))
	log.Printf("Named Pipe?   %v", (mode & os.ModeNamedPipe))
}
