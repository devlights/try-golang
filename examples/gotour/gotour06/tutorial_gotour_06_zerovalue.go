package gotour06

import (
	"fmt"
	"strings"
)

// ZeroValue は、 Tour of Go - Zero values (https://tour.golang.org/basics/12) の サンプルです。
// noinspection GoNilness
func ZeroValue() error {
	// ------------------------------------------------------------
	// 初期値
	// 変数に初期値を与えずに宣言すると、「ゼロ値」が設定される
	// ゼロ値は型によって異なる
	//
	// - 数値型は 0
	// - bool型は false
	// - string型は ""
	// - 配列は長さ0の配列
	// - スライスはnil
	// - マップはnil
	// ------------------------------------------------------------
	var (
		zeroInt    int
		zeroBool   bool
		zeroString string
		zeroArray  [0]int
		zeroSlice  []int
		zeroMap    map[int]int
	)

	_print(zeroInt, zeroBool, zeroString, zeroArray, zeroSlice, zeroMap)
	_print(len(zeroArray), len(zeroSlice), len(zeroMap))

	if zeroSlice == nil {
		fmt.Println("zeroSlice is nil")
	}

	if zeroMap == nil {
		fmt.Println("zeroMap is nil")
	}

	return nil
}

func _print(items ...interface{}) {
	formats := make([]string, len(items))
	for i := 0; i < len(formats); i++ {
		formats[i] = "%#v"
	}

	format := strings.Join(formats, "\t")
	if len(format) > 0 {
		format = fmt.Sprintf("%v\n", format)
	}

	fmt.Printf(format, items...)
}
