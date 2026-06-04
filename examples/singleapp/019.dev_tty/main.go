package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/term"
)

func main() {
	log.SetFlags(0)

	//
	// 以下、サンプルであるためエラー処理は割愛
	//

	// 通常の標準出力・標準エラー出力が生きているかの確認用
	fmt.Fprintln(os.Stdout, "stdout: hello")
	fmt.Fprintln(os.Stderr, "stderr: hello")

	// /dev/tty を読み書きモードで開く
	///
	// /dev/tty は「このプロセスの制御端末」へのエイリアスであり、
	// stdin/stdout/stderr のリダイレクト状態に関わらず
	// 常にユーザーのターミナルへ直接アクセスできる
	var (
		tty *os.File
	)
	tty, _ = os.OpenFile("/dev/tty", os.O_RDWR, 0)
	defer tty.Close()

	// /dev/tty 経由でターミナルに直接書き込む
	//
	// stdout が /dev/null に向いていても、これはターミナルに表示される
	fmt.Fprintln(tty, "tty: hello")

	// /dev/tty からの入力を行単位で読み込み、読み込んだ入力を各出力先(tty, stdout, stderr)に書き込む
	//
	// tty    : リダイレクトに関わらずターミナルに表示される
	// stdout : stdout が /dev/null の場合は捨てられる
	// stderr : stderr が /dev/null の場合は捨てられる
	var (
		reader *bufio.Reader
		line   string
	)
	fmt.Fprint(tty, "input> ")

	reader = bufio.NewReader(tty)
	line, _ = reader.ReadString('\n') // 末尾の改行ごと取得 (Scannerの場合は改行なしで取得)

	fmt.Fprintf(tty, "tty-echo   : %s", line)
	fmt.Fprintf(os.Stdout, "stdout-echo: %s", line)
	fmt.Fprintf(os.Stderr, "stderr-echo: %s", line)

	// エコーなしでパスワードを読み込む
	//
	// term.ReadPassword はシステムコールレベルで fd を直接操作するため
	// *os.File ではなく int の fd が必要。なお、Fd() を呼ぶとファイルがブロッキングモードに切り替わる。
	//
	// 内部で termios の ECHO フラグを一時的に無効化し、
	// 読み込み完了後に元の termios 設定が復元される
	var (
		fd = int(tty.Fd())
		pw []byte
	)
	fmt.Fprint(tty, "password> ")

	pw, _ = term.ReadPassword(fd)

	// term.ReadPassword はエコーを抑制するので改行も表示されない。
	// なので明示的に改行を出力して表示を調整。
	fmt.Fprintln(tty)

	fmt.Fprintf(tty, "tty-echo   : %s\n", pw)
	fmt.Fprintf(os.Stdout, "stdout-echo: %s\n", pw)
	fmt.Fprintf(os.Stderr, "stderr-echo: %s\n", pw)
}
