package zerovalues

import "github.com/devlights/gomy/output"

// Slice は、Goにおける スライス のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go-tour-jp.appspot.com/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Slice() error {
	//
	// スライス の ゼロ値 は nil
	// 注意点として スライス の場合、ゼロ値をprintfすると [] と表示される.
	// アドレスを表示すると ゼロ値 の場合は 0x0 と表示される.
	// また、append() には ゼロ値のスライス を渡すことが可能な点に注意.
	//
	var (
		s []int
	)

	output.Stdoutf("[slice zerovalue]", "%v\t%p\tNIL?=%t\n", s, s, s == nil)

	s = append(s, 100)
	output.Stdoutf("[slice after append]", "%v\t%p\tNIL?=%t\n", s, s, s == nil)

	return nil
}
