package variables

import "fmt"

// パッケージ スコープ な変数は普通に関数の外で変数定義すればよい
// var の 使い方は、var_statement_declares.go を参照。
var (
	// Go では、変数名の先頭を大文字にするか否かで変数の公開範囲が決まる
	// 先頭を大文字にすると public, 先頭が大文字以外の場合は private となる
	pkgScopeVal1 = "private package scope variable"
	PkgScopeVal1 = "public  package scope variable"
)

// PackageScopeVariable -- パッケージ スコープな変数のサンプル
func PackageScopeVariable() error {

	fmt.Printf("%#v\n%#v\n", pkgScopeVal1, PkgScopeVal1)

	return nil
}
