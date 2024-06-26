package packagescope

import "fmt"

// パッケージ スコープ な変数は普通に関数の外で変数定義すればよい
// var の 使い方は、var_statement_declares.go を参照。
var (
	// Go では、変数名の先頭を大文字にするか否かで変数の公開範囲が決まる
	// 先頭を大文字にすると public, 先頭が大文字以外の場合は private となる
	pkgScopeVal1 = "private package scope variable"
	PkgScopeVal1 = "public  package scope variable"
)

// Basic -- パッケージ スコープな変数のサンプル
func Basic() error {

	fmt.Printf("%#v\n%#v\n", pkgScopeVal1, PkgScopeVal1)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: package_scope_variable

	   [Name] "package_scope_variable"
	   "private package scope variable"
	   "public  package scope variable"


	   [Elapsed] 4.52µs
	*/

}
