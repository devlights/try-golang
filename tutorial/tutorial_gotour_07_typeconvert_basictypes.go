package tutorial

import "fmt"

// TypeConvertBasicTypes は、 Tour of Go - Zero values (https://tour.golang.org/basics/13) の サンプルです。
func TypeConvertBasicTypes() error {
	// ------------------------------------------------------------
	// 基本型の型変換
	// Go言語では、キャストしたい型を関数呼び出しのように指定してキャストする.
	// Go言語では、「暗黙の型変換」は存在しない。必ず明示的に変換を指定しないといけない
	// ------------------------------------------------------------
	var (
		i  = 100
		ui uint
		f  float32
	)

	// uint に変換
	ui = uint(i)
	// float に変換
	f = float32(i)

	fmt.Printf("%T\t%T\t%T\n%v\t%v\t%v\n", i, ui, f, i, ui, f)

	return nil
}
