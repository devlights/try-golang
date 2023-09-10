package zerovalues

import "github.com/devlights/gomy/output"

// Pointer は、Goにおける ポインタ のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Pointer() error {
	//
	// ポインタ の ゼロ値 は nil
	//
	var (
		p *int
	)

	output.Stdoutf("[pointer zerovalue]", "%v\n", p)

	return nil
}
