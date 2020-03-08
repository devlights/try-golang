package logging

import "log"

// Flags は、log.Flag()とlog.SetFlags()の挙動を確認するサンプルです.
func Flags() error {
	// ----------------------------------------------------------------
	// log.Flags()とlog.SetFlags()について
	//
	// log.Flags() は、現在設定されているログの設定フラグを取得することが出来る.
	// log.SetFlags()は、新たにログの設定フラグを設定することが出来る.
	//
	// デフォルトでは、log.LstdFlags が設定されている.
	// なのでそのまま、log.Println() とかを呼ぶと
	//   日付 時刻 メッセージ
	// という出力になる.
	//
	// この状態から、日付と時刻の出力を消したい場合は、log.LstdFlags を除去すれば良い.
	// フラグ定数は、$GOROOT/src/log/log.go の中に記載されている.
	//
	// REFERENCES::
	//   https://golang.org/pkg/log/#pkg-constants
	//   https://golang.org/pkg/log/#Flags
	//   https://golang.org/pkg/log/#SetFlags
	//   https://stackoverflow.com/questions/48629988/remove-timestamp-prefix-from-go-logger
	// ----------------------------------------------------------------
	// デフォルトの状態でログ出力
	log.Println("デフォルト状態でログ出力")

	// 現在 log.LstdFlags が設定されているかを確認
	log.Printf("log.LstdFlags [%t]", log.Flags()&log.LstdFlags == log.LstdFlags)

	// log.LstdFlags を 除去
	log.SetFlags(log.Flags() &^ log.LstdFlags)

	// 日付と時刻が出力されないかを確認
	log.Println("日付と時刻が消える")

	// 新たな フラグ として、ファイル名を表示するように設定
	log.SetFlags(log.Flags() | log.Lshortfile)

	// ファイル名が出力されるかを確認
	log.Println("ファイル名が付与される")
	log.Printf("log.Lshortfile [%v]", log.Flags()&log.Lshortfile == log.Lshortfile)

	// 一気に全設定フラグを落とす場合は以下のようにする
	log.SetFlags(0)

	// ただの文字列だけが出力されることを確認
	log.Println("全設定フラグをオフ")

	// 元に戻す
	log.SetFlags(log.Flags() | log.LstdFlags)

	// 元の状態に戻ったことを確認
	log.Println("元の状態に戻る")

	return nil
}
