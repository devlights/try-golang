package loops

import (
	"github.com/devlights/gomy/output"
)

// WhileLoop は、GoでのWhileループについてのサンプルです.
func WhileLoop() error {
	// Go には、ループはすべて for で記載することになっている。
	// 他の言語にある while () {} は提供されていない。
	count := 5
	for count > 0 {
		output.Stdoutl("[count]", count)
		count -= 1
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_while_loop

	   [Name] "loops_while_loop"
	   [count]              5
	   [count]              4
	   [count]              3
	   [count]              2
	   [count]              1


	   [Elapsed] 78.26µs
	*/

}
