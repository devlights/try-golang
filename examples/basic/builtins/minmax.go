package builtins

import "github.com/devlights/gomy/output"

// MinMax は、Go 1.21 で追加されたビルトイン関数のmin,maxについてのサンプルです.
//
// go doc で調べる場合は以下のようにします.
//
//   - go doc builtin.max
//   - go doc builtin.min
//
// # REFERENCES
//   - https://go.dev/ref/spec#Min_and_max
func MinMax() error {
	//
	// max, min は、Go 1.21から追加されたビルトイン関数
	//
	// - スライスは指定できない
	//
	var (
		x1 = 10
		x2 = 20
	)
	output.Stdoutf("[max(x1, x2)]", "%v\n", max(x1, x2))
	output.Stdoutf("[min(x1, x2)]", "%v\n", min(x1, x2))

	var (
		x3 = 2
	)
	output.Stdoutf("[max(x1, x2, x3)]", "%v\n", max(x1, x2, x3))
	output.Stdoutf("[min(x1, x2, x3)]", "%v\n", min(x1, x2, x3))

	// min, maxの型引数は cmp.Ordered となっているので 数値系と文字列 を渡せる
	var (
		x4 = "hello"
		x5 = "world"
	)
	output.Stdoutf("[max(x4, x5)]", "%v\n", max(x4, x5))
	output.Stdoutf("[min(x4, x5)]", "%v\n", min(x4, x5))

	// スライスと配列は渡すことが出来ない
	// min([]int{1,2,3})  // コンパイルエラー
	// min([3]int{1,2,3}) // コンパイルエラー

	return nil
}
