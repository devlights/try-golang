package zerovalues

import "github.com/devlights/gomy/output"

// Map は、Goにおける マップ のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go-tour-jp.appspot.com/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Map() error {
	//
	// マップ の ゼロ値 は nil
	// 注意点として マップ の場合、ゼロ値をprintfすると map[] と表示される.
	// アドレスを表示すると ゼロ値 の場合は 0x0 と表示される.
	// スライスとは違い ゼロ値 のマップにはキーを追加出来ない.
	//
	var (
		m map[int]string
	)

	output.Stdoutf("[map zerovalue]", "%v\t%p\tNIL?=%t\n", m, m, m == nil)

	// 以下のようにエラーとなる
	// m[100] = "apple"
	// >> panic: assignment to entry in nil map

	return nil
}
