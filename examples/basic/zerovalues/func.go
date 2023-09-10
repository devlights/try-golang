package zerovalues

import "github.com/devlights/gomy/output"

// Func は、Goにおける 関数 のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Func() error {
	//
	// 関数 の ゼロ値 は nil
	//
	var (
		fn1 func()
		fn2 = func() {}
	)

	output.Stdoutf("[func zerovalue]", "%p:%p\n", fn1, fn2)

	return nil
}
