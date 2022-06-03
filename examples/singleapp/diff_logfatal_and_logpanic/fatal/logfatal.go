package main

import (
	"fmt"
	"log"
	"os"
)

var (
	l = log.New(os.Stderr, ">>> ", 0)
)

func main() {
	// log.Fatal は、os.Exit(1) を呼び出すのでdeferが呼ばれない
	// log.Fatal は、OSに戻り値 1 を返す (linuxの場合)
	//
	// REFERENCES:
	//   - https://zenn.dev/spiegel/books/error-handling-in-golang/viewer/panics
	defer fmt.Println("call defer")
	l.Fatalln("call log.Fatal")
}
