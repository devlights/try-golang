package zerovalues

import "github.com/devlights/gomy/output"

// Int は、Goにおける int のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Int() error {
	//
	// int の ゼロ値 は 0
	//
	var (
		i int
	)

	output.Stdoutf("[int zerovalue]", "%d\n", i)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: zerovalues_int

	   [Name] "zerovalues_int"
	   [int zerovalue]      0


	   [Elapsed] 7.89µs
	*/

}
