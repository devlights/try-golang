package cmpop

import (
	"cmp"

	"github.com/devlights/gomy/output"
)

// Compare は、cmp.Compare[T cmp.Ordered]() のサンプルです。
//
// 戻り値は、他の言語とルールは同じで
//
//   - x が y より小さい場合は -1
//   - x と y が同じ場合は　　  0
//   - x が y より大きい場合は  1
//
// となる。比較対象として指定出来るのは cmp.Ordered となっている。
//
// cmp.Or と組合せることにより、ソート処理が書きやすくなる。
//
// # REFERENCES
//   - https://pkg.go.dev/cmp@go1.22.2#Compare
//   - https://pkg.go.dev/cmp@go1.22.2#Ordered
func Compare() error {

	output.Stdoutl("[compare1]", cmp.Compare("hello", "HELLO"))
	output.Stdoutl("[compare2]", cmp.Compare(100, 100))
	output.Stdoutl("[compare3]", cmp.Compare(99, 100))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: cmpop_compare

	   [Name] "cmpop_compare"
	   [compare1]           1
	   [compare2]           0
	   [compare3]           -1


	   [Elapsed] 42.1µs
	*/

}
