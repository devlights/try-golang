package ifs

import "fmt"

type (
	// Hello() を実装していない 構造体
	helloNotImpl struct{}

	// Hello() を実装している 構造体
	helloImpl struct{}
)

// Hello -- ダックタイピング対象の関数
func (helloImpl) Hello() string {
	return "hello world"
}

// DuckTyping -- Go で、インターフェースの仕組みを用いたPythonチックなダックタイピングのやり方のサンプルです.
func DuckTyping() error {
	// ----------------------------------------------------------------
	// GoでPythonのようにダックタイピングするTips
	//
	// Goでは、インターフェースをその場で定義することが可能
	// それを利用すると、予めインターフェースを定義しておき、それの実装を構造体に付与する
	// という過程を踏まなくても、実行時に判定することが出来る.
	//
	// Pythonなどの動的言語では、所謂ダックタイピングがよく利用されるが
	// Goでも似たようなことは一応出来る。
	//
	// 以下では、構造体を実装段階では、まだ特定のインターフェースを定義していない状態で
	// Hello()というメソッドを構造体に付与している。
	//
	// 判定する段階で、その場で、 Hello() string を持つインターフェースを定義して
	// 判定してしまい、実装を持っているものだけを通すようにしている。
	//
	// キモとなるのは、以下の判定のやり方。
	//
	// _, ok := notImpl.(interface { Hello() string　} )
	//
	// 通常、型検証する場合は notImpl.(特定のインターフェース) のようにするが
	// ここで、その場でインターフェースを定義している。
	// ----------------------------------------------------------------
	var (
		notImpl interface{} = helloNotImpl{}
		impl    interface{} = helloImpl{}
	)

	callHello(notImpl)
	callHello(impl)

	return nil
}

func callHello(v interface{}) {

	if i, ok := v.(interface{ Hello() string }); ok {
		fmt.Printf("%T は、 Hello() を実装している (%v)\n", v, i.Hello())
	} else {
		fmt.Printf("%T は、 Hello() を実装していない \n", v)
	}
}
