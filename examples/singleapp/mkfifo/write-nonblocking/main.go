//go:build linux

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sys/unix"
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
	// 名前付きパイプをノンブロッキングモードで開く
	//   明示的なノンブロッキングモードの指定は os.OpenFile() では行えないため
	//   golang.org/x/sys/unix を利用する
	//
	var (
		fd  int
		f   *os.File
		err error
	)

	log.Println("[Before] unix.Open(unix.O_WRONLY|unix.O_NONBLOCK)")

	// 名前付きパイプを書込みでノンブロッキングモードで開く場合に O_WRONLY を指定すると
	// 読み込み側が開かれていない場合 `ENXIO(no such device or address)` が発生する。
	//
	// 逆に読み込みでノンブロッキングモードで開く場合はここではエラーとならず
	// 実際に読み込む際に書込みが行われていない場合に即EOFが返ることになる。
	// ノンブロッキングリードする処理は ../read-nonbloking/ を参照のこと。
	//
	// 読み込みと書込みでエラーが発生する箇所が異なる点に注意。
	//
	// REFERENCES
	//   - https://qiita.com/seriru13/items/39ed2431dfd959ad512e
	for {
		fd, err = unix.Open(fname, unix.O_WRONLY|unix.O_NONBLOCK, 0666)
		if err != nil {
			var sysErr syscall.Errno
			if errors.As(err, &sysErr) && sysErr == unix.ENXIO {
				log.Printf("[ENXIO] %s", sysErr)
				<-time.After(200 * time.Millisecond)
				continue
			}

			return err
		}

		break
	}

	f = os.NewFile(uintptr(fd), fname)
	if f == nil {
		return fmt.Errorf("invalid file descriptor")
	}
	defer f.Close() // ここで f.Close() しているので、上で unix.Close(fd) は不要

	log.Println("[After ] unix.Open(unix.O_WRONLY|unix.O_NONBLOCK)")

	//
	// データを書き込み
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
		timeout = 1500 * time.Millisecond
		done    = make(chan struct{})
		wg      sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

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
		log.Println("timeout")
	}

	close(done)
	wg.Wait()

	return nil
}
