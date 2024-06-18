//go:build linux

package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"
)

var (
	fname string
)

func init() {
	log.SetFlags(log.Lmicroseconds)

	flag.StringVar(&fname, "fname", "", "FIFO file name")
	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// 名前付きパイプを開く
	//   os.OpenFile() にて、モード指定で os.ModeNamedPipe を指定する
	//   os.O_RDOONLYで開くと読み取り専用となるが、この場合書込みが発生するまで
	//   ブロックされる。これはUNIXの名前付きパイプの挙動に従った動作である。
	//
	var (
		f   *os.File
		err error
	)

	log.Println("[Before] os.OpenFile(os.O_RDOONLY)")

	// 対象となる名前付きパイプに書込みが発生していない場合、ここでブロックされる。
	f, err = os.OpenFile(fname, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Println("[After ] os.OpenFile(os.O_RDOONLY)")

	//
	// データを読み取り
	//
	type (
		data struct {
			value string
			err   error
		}
	)
	var (
		reader  = bufio.NewReader(f)
		lines   = make(chan data)
		timeout = time.Second
	)

	go func() {
		line, err := reader.ReadString('\n')
		lines <- data{line, err}
	}()

	select {
	case line := <-lines:
		if line.err != nil {
			return line.err
		}

		log.Println(line.value)
	case <-time.After(timeout):
		// 上記で記載した通り、os.O_RDOONLYのみで開いている場合
		// ブロッキングモードとなっているため、os.OpenFile()の呼び出しの
		// 時点でブロックされることとなる。つまり、実際にデータを読み出す
		// タイミングでは待たされることが無いため、このタイムアウトは通らない。
		log.Println("timeout")
	}

	return nil
}
