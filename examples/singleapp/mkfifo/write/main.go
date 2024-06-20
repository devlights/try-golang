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
	//   os.O_WRONLYで開くと書込み専用となるが、この場合読み込みが発生するまで
	//   ブロックされる。これはUNIXの名前付きパイプの挙動に従った動作である。
	//
	var (
		f   *os.File
		err error
	)

	log.Println("[Before] os.OpenFile(os.O_WRONLY)")

	// 対象となる名前付きパイプに読み込みが発生していない場合、ここでブロックされる。
	f, err = os.OpenFile(fname, os.O_WRONLY, os.ModeNamedPipe)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Println("[After ] os.OpenFile(os.O_WRONLY)")

	//
	// データを書込み
	//
	type (
		data struct {
			numWrites int
			err       error
		}
	)
	var (
		writer  = bufio.NewWriter(f)
		results = make(chan data)
		timeout = time.Second
	)

	go func() {
		n, err := writer.WriteString("helloworld\n")
		if err == nil {
			err = writer.Flush()
		}

		results <- data{n, err}
	}()

	select {
	case r := <-results:
		if r.err != nil {
			return r.err
		}

		log.Printf("Write %d byte(s)", r.numWrites)
	case <-time.After(timeout):
		// 上記で記載した通り、os.O_WRONLYのみで開いている場合
		// ブロッキングモードとなっているため、os.OpenFile()の呼び出しの
		// 時点でブロックされることとなる。つまり、実際にデータを書き出す
		// タイミングでは待たされることが無いため、このタイムアウトは通らない。
		log.Println("timeout")
	}

	return nil
}
