//go:build unix

package main

import (
	"flag"
	"log"
	"os"

	"golang.org/x/sys/unix"
)

var (
	fname string
)

func init() {
	log.SetFlags(0)

	flag.StringVar(&fname, "fname", "", "fifo file name")
	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	//
	// 名前付きパイプを作成
	//   mkfifoコマンドを実行したのと同じ意味となる
	//
	//   $ mkfifo fname -m0666
	//
	// [REFERENCES]
	//   - https://tldp.org/LDP/lpg/node15.html
	//   - https://linuxcommand.net/mkfifo/
	//
	err = unix.Mkfifo(fname, 0666)
	if err != nil {
		return err
	}

	if fi, err := os.Stat(fname); err == nil {
		if fi.Mode()&os.ModeNamedPipe != 0 {
			log.Println("Named Pipe was created.")
		}
	}

	return nil
}
