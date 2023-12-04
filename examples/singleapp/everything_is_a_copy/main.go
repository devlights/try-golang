package main

import (
	"github.com/devlights/gomy/output"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	//
	// Goにおける代入は「全てコピー」となる.
	//   - 構造体を返す関数の結果を代入すると、その「構造体のコピー」が行われる.
	//   - ポインタを返す関数の結果を代入すると、その「メモリアドレスのコピー」が行われる.
	//   - 構造体を受け取る関数に、構造体を渡すと、その「構造体のコピー」が行われる.
	//   - ポインタを受け取る関数に、ポインタを渡すと、その「メモリアドレスのコピー」が行われる.
	//
	// Goの for range の挙動も同じ理屈で動いている.
	//

	// 構造体を返す関数の結果を代入すると、その「構造体のコピー」が行われる
	v1 := NewValueType(100)

	// ポインタを返す関数の結果を代入すると、「メモリアドレスのコピー」が行われる
	v2 := NewPointerType(100)

	v1.Change() // レシーバーにはコピーが渡されるので内部で値を変えても元は変わっていない
	v2.Change() // レジーバーにはメモリアドレスのコピーが渡されるので値を変えると元も変わる

	output.Stdoutf("[v1]", "%v\n", v1)
	output.Stdoutf("[v2]", "%v\n", v2)

	// for range 文は、for i, v := range xx と記述するので、代入が行われる(:=の存在).
	// range は、対象のオブジェクトをコピーしてからループを始める.
	// この際のコピーのルールは上記の理屈となる.
	//
	// また、ループ変数にはループが始まる直前に一つの変数が確保され
	// その変数に繰り返しループごとの値が代入される.
	// (つまり、ループ変数のアドレスを取ると全部同じアドレス)
	//
	// この場合も代入が発生しているので同じ理屈となる
	// (この同じ変数に繰り返し代入される挙動は Go 1.22 で変わる予定 (GOEXPERIMENT loopvar))
	v3 := []interface{ Change() }{
		NewValueType(1), NewValueType(2), NewValueType(3),
		NewPointerType(1), NewPointerType(2), NewPointerType(3),
	}

	for _, v := range v3 {
		v.Change()
	}

	output.Stdoutf("[v3]", "%v\n", v3)

	return nil

	/*
	   $ task
	   task: [run] go run .
	   [v1]                 100
	   [v2]                 200
	   [v3]                 [1 2 3 101 102 103]
	*/

}
