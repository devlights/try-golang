package zerovalues

import "github.com/devlights/gomy/output"

// String は、Goにおける string のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func String() error {
	//
	// string の ゼロ値 は 空文字 ("")
	//
	var (
		s string
	)

	output.Stdoutf("[string zerovalue]", "%q\n", s)

	return nil
}
