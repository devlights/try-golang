package loops

import "github.com/devlights/gomy/output"

// SliceLoop は、スライスのループについてのサンプルです.
func SliceLoop() error {
	var (
		items = []string{
			"golang",
			"java",
			"dotnet",
			"python",
		}
	)

	// スライスの foreach は、インデックスと値 となる
	for i, v := range items {
		output.Stdoutf("", "[%d] %s\n", i, v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_slice_loop

	   [Name] "loops_slice_loop"
	   [0] golang
	   [1] java
	   [2] dotnet
	   [3] python


	   [Elapsed] 31.84µs
	*/

}
