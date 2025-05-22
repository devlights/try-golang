package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// 元の標準出力を退避させ、パイプのWriter側に差し替え
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		// 元に戻す。このプログラムはこのまま終了するので別にしなくて良いが習慣として。
		os.Stdout = old
	}()

	// コマンドライン引数
	os.Args = append(os.Args, "-v", "hello", "-v", "world", "-v", "へろー", "-v", "ワールド")

	var wg sync.WaitGroup
	wg.Add(1)

	// パイプはノンバッファリングなので非同期処理が必須
	go func() {
		defer wg.Done()
		defer w.Close()
		main()
	}()

	wg.Wait()

	// 出力内容を確認
	want := []byte("hello,world,へろー,ワールド\n")
	got, _ := io.ReadAll(r)
	if !bytes.Equal(want, got) {
		t.Errorf("want: %s\tgot: %s", want, got)
	}
}
