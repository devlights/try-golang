package filepaths

import (
	"path/filepath"

	"github.com/devlights/gomy/output"
)

// Clean は、filepath.Clean のサンプルです。
//
// filepath.Clean は、パス文字列を「意味は同じだけど最短の形」に正規化するだけの関数です。
// ファイルシステムは一切見ず、純粋に文字列処理だけを行います。
//
// filepath.Clean はざっくり以下のことをしてくれる。
//
//  1. 連続するセパレータを 1 個にする
//
//  2. カレント(".")を消す
//
//  3. パスの途中の .. を解決して、直前の要素とセットで消す
//
//  4. 絶対パスの先頭の /.. を / に変える（ルートより上には行かせない）
//
//  5. 空文字列の置換 (パスが空文字の場合はカレントディレクトリに変換)
//
// # REFERENCES
//   - https://pkg.go.dev/path/filepath@go1.26.3#Clean
func Clean() error {
	// 1. 連続するセパレータを 1 個にする
	v := filepath.Clean("foo/bar//baz//////file.txt")
	output.Stdoutl("[No.1]", v)

	// 2. カレント(".")を消す
	v = filepath.Clean("./dokokano/dir/./file.txt")
	output.Stdoutl("[No.2]", v)

	// 3. パスの途中の .. を解決して、直前の要素とセットで消す
	v = filepath.Clean("./dokokano/dir/../file.txt")
	output.Stdoutl("[No.3]", v)

	// 4. 絶対パスの先頭の /.. を / に変える（ルートより上には行かせない）
	v = filepath.Clean("/../dokokano/dir/file.txt")
	output.Stdoutl("[No.4]", v)

	// 5. 空文字列の置換 (パスが空文字の場合はカレントディレクトリに変換)
	v = filepath.Clean("")
	output.Stdoutl("[No.5]", v)

	return nil
}
