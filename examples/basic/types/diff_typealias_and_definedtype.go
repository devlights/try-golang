package types

import (
	"strings"

	"github.com/devlights/gomy/output"
)

type typeAlias = string // stringの別名を定義（stringとしてそのまま利用可能)
type definedType string // ベースが string の全く別の型を定義 (string(v)とすることで変換可能)

// NG コンパイルエラー
//   invalid receiver type string (basic or unnamed type)
// func (me aliasType) Upper() aliasType {
// }

// OK
//lint:ignore U1000 ok
func (me definedType) upper() string {
	return strings.ToUpper(string(me))
}

// DiffTypeAliasAndDefinedType -- Goでの Type Alias と Defined Type の違いについてのサンプルです.
//
// 個人的に、Type Alias はほぼ使ったことがない。基本的に Defined Type で事足りている。
//
// REFERENCES
//   - https://yourbasic.org/golang/type-alias/
//   - https://budougumi0617.github.io/2020/02/01/go-named-type-and-type-alias/
//   - https://text.baldanders.info/golang/go-1_9-and-type-alias/
func DiffTypeAliasAndDefinedType() error {
	var (
		original string = "hello world"
	)

	// ----------------------------------------------
	// ベースとなる string の値を設定する場合の違い
	// ----------------------------------------------

	// OK
	var alias typeAlias = original

	// NG
	//   cannot use original (variable of type string) as definedType value in variable declaration
	// var defined definedType = original

	// OK
	var defined definedType = definedType(original)

	// ----------------------------------------------
	// string の引数を要求する関数に渡す場合の違い
	// ----------------------------------------------

	// OK
	show(alias)

	// NG
	//   cannot use defined (variable of type definedType) as string value in argument to show
	// show(defined)

	// OK
	show(string(defined))

	return nil
}

func show(v string) {
	output.Stdoutf("[v]", "%v (%T)\n", v, v)
}
