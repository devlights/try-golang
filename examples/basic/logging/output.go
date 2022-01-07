package logging

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

// Output -- Writer は、 log.Writer() と log.SetOutput() の挙動を確認するサンプルです.
func Output() error {
	// ----------------------------------------------------------------
	// log.Writer() と log.SetOutput() について
	//
	// log.Writer() は、標準ロガーに現在設定されている io.Writer を取得することができる
	// log.SetWriter() は、標準ロガーに新たなログの出力先である io.Writer を設定することができる
	//
	// log.Writer() のデフォルトは 標準エラー出力. (os.Stderr)
	// 標準ロガーは、$GOROOT/src/log/log.go で以下のように定義されている
	//   var std = New(os.Stderr, "", LstdFlags)
	// log.Writer() の実装は、std.Writer を返しているだけ.
	// ----------------------------------------------------------------
	origFlags := log.Flags()
	origWriter := log.Writer()
	log.SetFlags(0)
	defer log.SetFlags(origFlags)
	defer log.SetOutput(origWriter)

	// この時点では、まだ 出力先 を変更していないので標準エラーに出力される
	log.Println("SetOutput 呼び出す前")

	// 標準ロガーの出力先を bytes.Buffer に切り替え
	var buf bytes.Buffer
	log.SetOutput(&buf)

	_, _ = fmt.Fprintf(os.Stderr, "buf == log.Writer() [%v]\n", &buf == log.Writer())

	// このメッセージは、標準エラーには出力されずにバッファに送られる
	log.Println("SetOutput 呼び出し後")

	// バッファから内容を取り出して、出力
	s := buf.String()
	_, _ = fmt.Fprintf(os.Stderr, "[fmt.Printf] %s", s)

	return nil
}
