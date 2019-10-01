package tutorial

import "fmt"

// パッケージレベルでの定数宣言
const packageScopeConstants = "hello"

func GoTourConstant() error {
	// ------------------------------------------------------------
	// Go言語では、定数は const キーワードを用いて宣言する
	// 定数は、文字、文字列、bool値,数値でのみ利用できる
	// 定数は、 := を使用して宣言はできない
	// ------------------------------------------------------------
	const localScopeConstants = "world"

	// 当然ながら、定数に再代入は不可能
	// (cannot assign to localScopeConstants)
	// localScopeConstants = "hoge"

	fmt.Println(packageScopeConstants, localScopeConstants)

	return nil
}
