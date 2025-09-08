package sliceop

import (
	"github.com/devlights/gomy/output"
)

// Concat -- ２つのスライスの結合に関するサンプルです.
func Concat() error {
	// -------------------------------------------------------------
	// スライスの結合
	//
	// (1) for ループで回して結合スライス作る
	// (2) append で結合スライス作る
	// -------------------------------------------------------------
	var (
		sli1 = []int{1, 2, 3, 4, 5}
		sli2 = []int{6, 7, 8}
	)

	// (1)
	sli3 := make([]int, 0, len(sli1)+len(sli2))
	for _, s := range [][]int{sli1, sli2} {
		sli3 = append(sli3, s...)
	}

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli2]", sli2)
	output.Stdoutl("[sli3]", sli3)
	output.StdoutHr()

	// (2)
	sli4 := append(sli1, sli2...)

	output.Stdoutl("[sli1]", sli1)
	output.Stdoutl("[sli2]", sli2)
	output.Stdoutl("[sli4]", sli4)
	output.StdoutHr()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_concat

	   [Name] "slice_concat"
	   [sli1]               [1 2 3 4 5]
	   [sli2]               [6 7 8]
	   [sli3]               [1 2 3 4 5 6 7 8]
	   --------------------------------------------------
	   [sli1]               [1 2 3 4 5]
	   [sli2]               [6 7 8]
	   [sli4]               [1 2 3 4 5 6 7 8]
	   --------------------------------------------------


	   [Elapsed] 59.66µs
	*/

}
