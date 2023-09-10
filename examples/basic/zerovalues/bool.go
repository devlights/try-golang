package zerovalues

import "github.com/devlights/gomy/output"

// Bool は、Goにおける bool のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Bool() error {
	//
	// bool の ゼロ値 は false
	//
	var (
		b bool
	)

	output.Stdoutf("[bool zerovalue]", "%t\n", b)

	return nil
}
