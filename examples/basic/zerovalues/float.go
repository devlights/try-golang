package zerovalues

import "github.com/devlights/gomy/output"

// Float は、Goにおける float のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: zerovalues_float

	   [Name] "zerovalues_float"
	   [float zerovalue]    0.0


	   [Elapsed] 11.64µs
	*/

}
