package zerovalues

import "github.com/devlights/gomy/output"

// Float は、Goにおける float のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go-tour-jp.appspot.com/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Float() error {
	//
	// float (float32, float64) の ゼロ値 は 0.0
	//
	var (
		f float64
	)

	output.Stdoutf("[float zerovalue]", "%.1f\n", f)

	return nil
}
