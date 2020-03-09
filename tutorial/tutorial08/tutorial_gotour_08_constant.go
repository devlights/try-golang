package tutorial08

import "fmt"

const (
	// パブリックスコープな定数宣言 (先頭を大文字で開始)
	PublicScopeConstants = "helloworld"

	// パッケージレベルでの定数宣言 (先頭を小文字で開始)
	packageScopeConstants = "hello"
)

// Constant は、 Tour of Go - Constant (https://tour.golang.org/basics/15) の サンプルです。
func Constant() error {
	// ------------------------------------------------------------
	// Go言語では、定数は const キーワードを用いて宣言する
	// 定数は、文字、文字列、bool値,数値でのみ利用できる
	// 定数は、 := を使用して宣言はできない
	// ------------------------------------------------------------
	const localScopeConstants = "world"

	// 当然ながら、定数に再代入は不可能
	// (cannot assign to localScopeConstants)
	// localScopeConstants = "hoge"

	fmt.Println(PublicScopeConstants, packageScopeConstants, localScopeConstants)

	return nil
}
