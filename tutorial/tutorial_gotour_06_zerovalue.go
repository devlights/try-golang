package tutorial

import "fmt"

func GoTourZeroValue() error {
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
