package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	ErrIoTimeout = errors.New("タイムアウト発生")
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.SetOutput(os.Stderr)

	var (
		rootCtx  = context.Background()
		ctx, cxl = context.WithTimeout(rootCtx, 5*time.Second)
		err      error
	)
	defer cxl()

	log.Println("start")
	if err = run(ctx); err != nil {
		if errors.Is(err, ErrIoTimeout) {
			log.Println(err)
			return
		}

		panic(err)
	}
}

func run(pCtx context.Context) error {
	//
	// os.FileにはSetDeadlineメソッドがあり、設定することでタイムアウトを仕込むことが出来る。
	// しかし、このメソッドは通常ファイルの場合は何も起こらない。（タイムアウトが発生しない）
	// 理由は通常ファイルがpoll可能なFDでは無いため。
	//
	// Goは内部で netpoller (Linuxの場合は epoll, macOSの場合は kqueue) を用いて非同期I/Oを実現している。
	//
	// - https://go.dev/src/runtime/netpoll.go
	// - https://morsmachine.dk/netpoller.html
	// - https://internals-for-interns.com/posts/go-netpoller/
	//
	// なので、poll可能なFDしかタイムアウトを設定出来ない。
	//
	// ソケット　　： 可能
	// パイプ　　　： 可能
	// TTY 　　　　： 可能
	// 通常ファイル： 不可
	//
	// 通常ファイルは poll 出来ないFDなので、SetDeadlineメソッドの呼び出しは出来るが何も起こらない。
	// いろいろなやり方があるが、パイプを使ってデータを流し、タイムアウトを設定するというやり方もある。
	//

	// (1) まずファイルを普通に開く
	const (
		fpath = "main.go" // このファイル
	)
	var (
		file *os.File
		err  error
	)
	if file, err = os.Open(fpath); err != nil {
		return fmt.Errorf("os.Open(%s): %w", fpath, err)
	}
	defer file.Close()

	// (2) パイプを生成
	var (
		pr *os.File
		pw *os.File
	)
	if pr, pw, err = os.Pipe(); err != nil {
		return fmt.Errorf("os.Pipe(): %w", err)
	}
	defer pr.Close()

	// (3) ファイルデータをパイプに流す
	//     w.Close() の呼び出しで pr 側にEOFが返る
	//
	// 実際に実行すると、本サンプルの場合は io.Copy は即座に完了する。
	// これは パイプ が、カーネル空間のバッファを利用しており、そのデフォルト値に収まっているから。
	//
	// 通常のLinuxの場合、Linuxカーネルのパイプバッファ初期値は実際には16KB (4096バイトが4ページ分)となっている。
	// この値を超えるデータが流れている場合は、当然 pr 側が読み取りを行うまでブロックされる。
	//
	// pythonで確認するのが楽
	//   import os,fcntl
	//   r,w = os.pipe()
	//   print(fcntl.fcntl(w, 1032)) # 1032 == F_GETPIPE_SZ
	go func(r io.Reader, w io.WriteCloser) {
		defer w.Close()
		io.Copy(w, r)
		log.Println("pw: io.Copy()")
	}(file, pw)

	// (4) パイプから読み取りながら、所定時間後にタイムアウト発生させる
	//
	// サンプルなので500msごとに１文字ずつ読み取っていき確実にタイムアウトするようにしている
	var (
		timeout  = 3 * time.Second
		interval = 500 * time.Millisecond
		tick     = time.NewTicker(interval)
		buf      = make([]byte, 1)
	)
	defer tick.Stop()

	// タイムアウト設定
	//   パイプは poll 可能なので SetDeadline が有効
	//   SetDeadline/SetReadDeadline/SetWriteDeadlineに現在時刻を設定することで即座にタイムアウトとなる
	go func(f *os.File, timeout time.Duration) {
		time.Sleep(timeout)
		f.SetDeadline(time.Now())
		log.Println("pr: SetDeadline()")
	}(pr, timeout)

	for {
		clear(buf)

		select {
		case <-pCtx.Done():
			return pCtx.Err()
		case <-tick.C:
			if _, err = pr.Read(buf); err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}

				if errors.Is(err, os.ErrDeadlineExceeded) {
					// タイムアウト
					return fmt.Errorf("%w (%w)", ErrIoTimeout, err)
				}

				return err
			}

			os.Stdout.Write(buf)
		}
	}
}
