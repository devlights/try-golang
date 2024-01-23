package loops

import "github.com/devlights/gomy/output"

// BasicForeach は、Go での foreach ループについてのサンプルです.
func BasicForeach() error {
	var (
		items = []string{
			"go",
			"java",
			"dotnet",
			"python",
			"flutter",
		}
	)

	for i, v := range items {
		output.Stdoutf("", "[%d] %s\n", i, v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_basic_foreach

	   [Name] "loops_basic_foreach"
	   [0] go
	   [1] java
	   [2] dotnet
	   [3] python
	   [4] flutter


	   [Elapsed] 71.68µs
	*/

}
