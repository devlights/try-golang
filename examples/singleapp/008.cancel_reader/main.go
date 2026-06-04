package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/muesli/cancelreader"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.SetOutput(os.Stdout)

	if err := run(); err != nil {
		log.Panic(err)
	}
}

func run() error {
	//
	// 標準入力をキャンセル可能なReaderにする
	//
	// cancelreader.CancelReaderは、Linuxの場合
	// 内部で epoll を利用して制御を行っている。
	//
	// なので、非ブロッキングI/Oが可能なファイルディスクリプタに対してのみ
	// キャンセル可能となっている。非ブロッキングI/Oが可能なものは例えば以下のもの。
	//   - ソケット
	//   - パイプ
	//   - FIFO
	//
	// 通常のファイルは非ブロッキングI/O可能なファイルディスクリプタでは無いことに注意。
	// 通常のファイルをepollで利用すると epoll_ctl() で EPERM が返ってくる。
	// 通常のファイルをキャンセル可能にしたい場合は io_uring などを検討する。
	//

	var (
		reader cancelreader.CancelReader
		err    error
	)
	reader, err = cancelreader.NewReader(os.Stdin)
	if err != nil {
		return fmt.Errorf("cancelreader.NewReader() failed: %w", err)
	}
	defer reader.Close()

	//
	// ３秒後にキャンセル
	//
	go func() {
		time.Sleep(3 * time.Second)
		reader.Cancel()
	}()

	//
	// 500ms毎に１文字読み込み
	//
	var (
		buf [1]byte
	)
	for {
		clear(buf[:])

		if _, err = reader.Read(buf[:]); err != nil {
			if errors.Is(err, cancelreader.ErrCanceled) {
				log.Println("CANCELED")
				break
			}

			return fmt.Errorf("reader.Read() failed: %w", err)
		}

		log.Print(string(buf[:]))
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}
