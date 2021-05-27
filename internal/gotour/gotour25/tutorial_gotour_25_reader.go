package gotour25

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Reader は、 Tour of Go - Readers (https://tour.golang.org/methods/21) の サンプルです。
func Reader() error {
	// ------------------------------------------------------------
	// Go言語の io.Reader インターフェース
	// io.Readerインターフェースは、何かを読み取る基本的なインターフェース
	// このインターフェースを実装している型はとても多い。
	//
	// io.Reader インターフェースには、以下の定義がある
	//   - Read(b []byte) (n int, err error)
	//
	// 基本的なものとして、 strings.Reader がある
	// これは、他の言語の StringReader と同じような使い方が出来る
	//
	// n に読み込んだバイト数が設定されるので、それでバッファから取り出して
	// 末尾まで読み込むと、 err に io.EOF が設定されるので
	// それを判断して処理を中断するというのが基本パターン.
	// ------------------------------------------------------------
	var (
		message   = "hello world"
		reader    = strings.NewReader(message)
		chunkSize = 4
		buf       = make([]byte, chunkSize)
		results   = make([]byte, 0, len(message))
	)

	// 4 バイトずつ読み込み
	for {
		n, err := reader.Read(buf)
		fmt.Printf("[reader] read: %dbytes\terror: %v\tvalue: %s\n", n, err, buf[:n])

		if n != 0 {
			results = append(results, buf[:n]...)
		}

		if errors.Is(err, io.EOF) {
			break
		}
	}

	fmt.Printf("[result] %s\n", string(results))

	// (補足) 上の例ではRead用のバッファと結果用のバッファの２つを使っているが
	//        一つのバッファで読み込み続けることも可能。２つバッファ使った方が
	//        わかりやすいので個人的にはあまり使わない。
	var (
		from      = 0
		to        = 0
		readBytes = 0
	)

	// 1 バイトずつ読み込み
	reader = strings.NewReader(message)
	buf = make([]byte, len(message))
	for {
		from, to = readBytes, readBytes+1
		if len(message) < to {
			break
		}

		n, err := reader.Read(buf[from:to])
		readBytes += n

		if n == 0 || err != nil {
			break
		}
	}

	fmt.Printf("[result] %d bytes. (%s)\n", readBytes, string(buf))

	return nil
}
