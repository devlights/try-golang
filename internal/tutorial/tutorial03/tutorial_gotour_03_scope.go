package tutorial03

import "fmt"

// Scope は、 Tour of Go - Exported names (https://tour.golang.org/basics/3) の サンプルです。
func Scope() error {
	// ------------------------------------------------------------
	// Goでは、最初の文字が大文字で始まる名前は外部に公開される。(public扱い)
	// 小文字で始まる名前は外部に公開されない。(private扱い)
	// 公開範囲は、パッケージ単位。なので、小文字の名前をつけた要素も同一パッケージ内
	// では、見ることが出来る。
	// ------------------------------------------------------------
	fmt.Println(packagePrivateFunc())

	return nil
}

// 小文字で始まっているのでこの関数は非公開関数 (パッケージプライベート)
func packagePrivateFunc() string {
	return "This is package-private function"
}
