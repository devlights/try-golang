package gotour07

import "fmt"

type (
	int2 int
)

// TypeConvertBasicTypes は、 Tour of Go - Type conversions (https://tour.golang.org/basics/13) の サンプルです。
func TypeConvertBasicTypes() error {
	// ------------------------------------------------------------
	// 基本型の型変換
	// Go言語では、キャストしたい型を関数呼び出しのように指定してキャストする.
	// Go言語では、「暗黙の型変換」は存在しない。必ず明示的に変換を指定しないといけない
	//
	// type int2 int と別名を付けた場合でも同じ。明示的なキャストが必要。
	// ------------------------------------------------------------
	var (
		i  = 100
		i2 int2
		ui uint
		f  float32
	)

	// uint に変換
	ui = uint(i)
	// float に変換
	f = float32(i)
	// int2 に変換
	i2 = int2(i)

	fmt.Printf("%T\t%T\t%T\t%T\n%v\t%v\t%v\t%v\n", i, i2, ui, f, i, i2, ui, f)

	return nil
}
