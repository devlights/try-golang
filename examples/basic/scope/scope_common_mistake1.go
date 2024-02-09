package scope

import (
	"os"

	"github.com/devlights/gomy/output"
)

var (
	// 現在の作業ディレクトリ. CommonMistake関数で使う前に初期化されている想定
	_cwd1 string
)

func loadcwd() error {
	// 一見、ちゃんと os.Getwd() の結果を パッケージ変数 _cwd1 に設定できているように見えるが・・
	_cwd1, err := os.Getwd()
	if err != nil {
		return err
	}

	// ここでの結果はちゃんと表示される
	output.Stdoutl("[loadcwd]", _cwd1)

	return nil
}

// CommonMistake1 -- 変数宣言のスコープによるよくやる間違いについてのサンプルです.
func CommonMistake1() error {
	if err := loadcwd(); err != nil {
		return err
	}

	output.Stdoutl("[main]", _cwd1)

	// -------------------------------------------------------------------------
	// 実行すると、上の _cwd1 の値は空文字で出力される.
	// 理由は、loadcwd() で設定している _cwd1 は、省略変数宣言 := を使っているため
	// ローカル変数 _cwd1 が新たに生成されてしまったため。
	//
	// loadcwd() 内で、最後にログ出力変わりに ローカル変数 _cwd1 の値を出力する
	// ようにしているため、一見ちゃんと設定できているように見えるし
	// ログ出力するために変数を使用しているため、コンパイルエラーにもならない
	// (これがログ出力部分がなかったら、ローカル変数を使用していないという
	//  コンパイルエラーとなるため、そこで気づける可能性がある。)
	// -------------------------------------------------------------------------

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: scope_common_mistake1

	   [Name] "scope_common_mistake1"
	   [loadcwd]            /workspace/try-golang
	   [main]

	   [Elapsed] 30.79µs
	*/

}
