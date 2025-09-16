//go:build unix

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
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

	log.Println("[Before] unix.Open(unix.O_RDONLY|unix.O_NONBLOCK)")

	fd, err = unix.Open(fname, unix.O_RDONLY|unix.O_NONBLOCK, 0666)
	if err != nil {
		return err
	}

	f = os.NewFile(uintptr(fd), fname)
	if f == nil {
		return fmt.Errorf("invalid file descriptor")
	}
	defer f.Close() // ここで f.Close() しているので、上で unix.Close(fd) は不要

	log.Println("[After ] unix.Open(unix.O_RDONLY|unix.O_NONBLOCK)")

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
		timeout = 1500 * time.Millisecond
		done    = make(chan struct{})
		wg      sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		// ノンブロッキングモードで処理しているため、データが存在しない場合は即EOFが返ってくる
		for {
			line, err := reader.ReadString('\n')
			if err != nil && errors.Is(err, io.EOF) {
				log.Println("読み取れるデータが存在しない")

				select {
				case <-done:
					return
				case <-time.After(200 * time.Millisecond):
					continue
				}
			}

			lines <- data{line, err}
			return
		}
	}()

	select {
	case line := <-lines:
		if line.err != nil {
			return line.err
		}

		log.Println(line.value)
	case <-time.After(timeout):
		log.Println("timeout")
	}

	close(done)
	wg.Wait()

	return nil
}
