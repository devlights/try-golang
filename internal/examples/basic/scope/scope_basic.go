package scope

import (
	"fmt"

	"github.com/devlights/try-golang/internal/examples/basic/scope/mypkg"
)

// Basic は、スコープについての基本的な事項についてのサンプルです.
//noinspection GoNameStartsWithPackageName
func Basic() error {
	// Go言語では、「最初の文字が大文字で始まる名前」は外部から参照可能な名前となる。
	// 他の言語でいうと public なスコープとなる。
	// 「小文字で始まる名前」は外部から参照不可な名前となる。
	// 他の言語でいうと private なスコープとなる。
	hasPublicFields := mypkg.HasPublicFields{}
	hasNoPublicFields := mypkg.HasNoPublicFields{}

	// こちらの構造体は 大文字で始まる名前 を持っているので参照できる
	hasPublicFields.Val1 = 100
	hasPublicFields.Val2 = "Test Value"

	// こちらの構造体は 小文字で始まる名前 を持っているので参照できない
	// コンパイルエラー
	// .\printf02.go:13:19: hasNoPublicFields.val1 undefined (cannot refer to unexported field or method val1)
	// hasNoPublicFields.val1 = 100

	ShowValues(&hasPublicFields)
	ShowValues(&hasNoPublicFields)

	return nil
}

// ShowValues -- サンプル関数
func ShowValues(obj mypkg.ICanDisplayValues) {
	fmt.Printf("[%T] %s\n", obj, obj.GetValues())
}
