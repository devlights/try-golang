package effectivego

import (
	"fmt"
	"log"
)

type (
	// effectivego10st01 構造体用のオプション定義 (Functional Options Pattern)
	effectivego10st01option func(*effectivego10st01) error

	// 本サンプル用の構造体
	effectivego10st01 struct {
		intVal    int
		boolVal   bool
		stringVal string
	}

	effectivego10st02 struct {
		IntVal    int
		BoolVal   bool
		StringVal string
	}
)

// effectivego10st01 用のコンストラクタ関数
func newEffectivego10st01(opts ...effectivego10st01option) *effectivego10st01 {
	v := new(effectivego10st01)
	v.intVal = -1
	v.boolVal = true
	v.stringVal = "hello world"

	// オプションが付与されていたら適用 (Functional Options パターン)
	for _, opt := range opts {
		if err := opt(v); err != nil {
			log.Fatal(err)
		}
	}

	// C と違い、 Go ではローカルで生成したポインタを呼び元に返しても問題にはならない。
	// 値は関数呼び出しが終わった後でも、ちゃんと紐付いている。
	// (Cでは関数内で宣言したポインタを呼び元に返すと、関数の呼び出し終了とともに解放されてしまうので無効なアドレスを操作することになる）
	return v
}

// effectivego10st01.intVal 用の オプション関数
func withIntVal(x int) effectivego10st01option {
	return func(e *effectivego10st01) error {
		e.intVal = x
		return nil
	}
}

// effectivego10st01.stringVal 用の オプション関数
func withStrVal(s string) effectivego10st01option {
	return func(e *effectivego10st01) error {
		e.stringVal = s
		return nil
	}
}

// Constructors -- Effective Go - Constructors and composite literals の 内容についてのサンプルです。
func Constructors() error {
	/*
			https://golang.org/doc/effective_go.html#composite_literals

			- Go には、他の言語にあるコンストラクタの仕組みが存在しない。
			  - なので、必要であれば自分で関数を定義して作る。
			- 基本的に、NewXXXX のような形で関数を作ることが多い模様
			- composite literals で初期化してしまうのも便利。
			- オプションを渡せるようにしたい場合は「Functional Options パターン」が便利
		      - https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
			  - https://qiita.com/weloan/items/56f1c7792088b5ede136
	*/
	// コンストラクタ関数を利用
	v1 := newEffectivego10st01()
	v2 := newEffectivego10st01(
		withIntVal(10),
		withStrVal("world hello"),
	)

	// Composite literalsを利用
	v3 := effectivego10st02{
		IntVal:    100,
		BoolVal:   true,
		StringVal: "hello world",
	}

	fmt.Printf("v1:%p (%v)\tv2:%p (%v)\tv3:%p (%v)\n", v1, *v1, v2, *v2, &v3, v3)

	return nil
}
