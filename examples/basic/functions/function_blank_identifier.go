package functions

var (
	_x = 1
	_y = 2 - _x
)

func _() {
	// この中のコードはコンパイル時にのみチェックされ、コンパイル結果には含まれない
	var x [1]struct{}
	_ = x[_x-1]
	_ = x[_y-1]
}

// BlankIdentifier は、 func _(){} という特殊関数を定義した場合の挙動についてのサンプルです.
//
// func _() { ... } という「関数名がアンダースコアになっている関数」は、特殊な目的で利用される。
//
// Goにおいて _ は、「無視する」や「使用しない」を意味するため、この関数は実際には利用されない。
// しかし、コンパイラはこの関数をコンパイル時にチェックする。
// なので、func _() {} の部分はコンパイル時に特定のチェックを行い、結果がNGの場合はコンパイルエラーに
// したい場合のコードを仕込む事ができる。
//
// 実際に、Stringerにてコードを生成されたものの中に上記の func _(){} が出力されている。
// Stringerの出力では、生成したENUMの元となる定数宣言が後から変更されている場合にコンパイルエラーと
// なるチェックが仕込まれている。
//
// 以下のようなものとなる。
//
//	func _() {
//		// An "invalid array index" compiler error signifies that the constant values have changed.
//		// Re-run the stringer command to generate them again.
//		var x [1]struct{}
//		_ = x[Placebo-0]
//		_ = x[Aspirin-1]
//		_ = x[Ibuprofen-2]
//	}
//
// このコードはコンパイル時にのみ有効であり、コンパイル結果には含まれない。
//
// # REFERENCES
//   - https://tech.uzabase.com/entry/2022/08/18/120148
func BlankIdentifier() error {
	//
	// func _(){} のサンプルであるため、本関数には実処理は無い
	//
	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: function_blank_identifier

	   [Name] "function_blank_identifier"


	   [Elapsed] 1.2µs
	*/

}
