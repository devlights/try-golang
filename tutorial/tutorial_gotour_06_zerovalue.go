package tutorial

import "fmt"

// ZeroValue は、 Tour of Go - Zero values (https://tour.golang.org/basics/12) の サンプルです。
func ZeroValue() error {
	// ------------------------------------------------------------
	// 初期値
	// 変数に初期値を与えずに宣言すると、「ゼロ値」が設定される
	// ゼロ値は型によって異なる
	//
	// - 数値型は 0
	// - bool型は false
	// - string型は ""
	// ------------------------------------------------------------
	var (
		zeroInt    int
		zeroBool   bool
		zeroString string
	)

	fmt.Printf("%v\t%v\t%q\n", zeroInt, zeroBool, zeroString)

	return nil
}
