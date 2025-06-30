package cmdexec

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// OsPipe は、(*Cmd).Stdout に os.Pipe の io.Writer を接続して処理するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.24.4#Pipe
func OsPipe() error {
	//
	// (*Cmd).StdoutPipe()で同じことが出来るが
	// os.Pipe()の使い方を勉強するために
	// 意図的に利用している
	//

	///////////////////////////////////
	// パイプ取得
	///////////////////////////////////

	var (
		pr  *os.File
		pw  *os.File
		err error
	)
	if pr, pw, err = os.Pipe(); err != nil {
		return err
	}
	defer pr.Close()

	///////////////////////////////////
	// コマンド実行
	//
	// git log コマンドを実行しているため
	// リポジトリによっては長大な出力が発生する。
	//
	// 簡易なコマンド実行である (*Cmd).Output() で取得しようとすると
	// OSのバッファが一杯になってしまう可能性があるため、このような場合は
	// ストリーミング処理が必須となる。
	//
	// 以下で実行している git コマンドのオプションは以下の通り
	//   - --no-pager  : ページャーを使用しない
	//   - log         : ログを表示
	//   - -m          : マージコミットの差分も表示
	//   - -r          : 再帰的に処理
	//   - --name-only : ファイル名のみ表示
	//   - --pretty=raw: 生フォーマットで表示
	//   - -z          : NULL文字で区切る
	///////////////////////////////////

	var (
		name = "git"
		args = []string{"--no-pager", "log", "-m", "-r", "--name-only", "--pretty=raw", "-z"}
		cmd  = exec.Command(name, args...)
	)
	cmd.Stdout = pw
	if err = cmd.Start(); err != nil {
		pw.Close()
		return err
	}

	///////////////////////////////////
	// 終了待機用のゴルーチンを用意
	///////////////////////////////////

	var (
		done = make(chan error, 1)
	)
	go func() {
		defer pw.Close()
		done <- cmd.Wait()
	}()

	///////////////////////////////////
	// コマンドの出力を読み出し
	///////////////////////////////////

	const (
		MaxTokenSize = 1024 * 1024 // 1行のサイズが大きい可能性を考慮してバッファサイズを底上げ
	)
	var (
		scanner = bufio.NewScanner(pr)
		buf     = make([]byte, MaxTokenSize)
		count   int
	)
	scanner.Buffer(buf, MaxTokenSize)

	for scanner.Scan() {
		io.Discard.Write(scanner.Bytes())
		count++
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	///////////////////////////////////
	// コマンド終了待機
	///////////////////////////////////

	if err = <-done; err != nil {
		return err
	}

	///////////////////////////////////
	// 結果出力
	///////////////////////////////////

	fmt.Printf("Total lines: %d\n", count)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmdexec_ospipe

	   [Name] "cmdexec_ospipe"
	   Total lines: 29988

	   [Elapsed] 123.966649ms
	*/
}
