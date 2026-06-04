package main

import (
	"log"
	"os"
)

var (
	l = log.New(os.Stderr, ">>> ", 0)
)

func main() {
	// log.Panic は、panicを呼び出すのでdeferが呼ばれる
	// log.Panic は、OSに戻り値 2 を返す (linuxの場合)
	//
	// REFERENCES:
	//   - https://zenn.dev/spiegel/books/error-handling-in-golang/viewer/panics
	defer l.Println("call defer")
	l.Panicln("call log.Panic")
}
