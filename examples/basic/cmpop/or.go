package cmpop

import (
	"cmp"

	"github.com/devlights/gomy/output"
)

// Or は、cmp.Or[T comparable]() のサンプルです。
// cmp.Or は、Go 1.22 で追加されました。
//
// > Or returns the first of its arguments that is not equal to the zero value. If no argument is non-zero, it returns the zero value.
//
// > (Orは、引数のうちゼロ値ではない最初の引数を返す。どの引数もゼロ値ではない場合、ゼロ値を返す。)
//
// comparableが対象となるので、以下に対して利用できる。
//
//   - booleans
//   - numbers
//   - strings
//   - pointers
//   - channels
//   - arrays of comparable types
//   - structs whose fields are all comparable types
//
// # REFERENCES
//   - https://pkg.go.dev/cmp#Or
//   - https://pkg.go.dev/builtin#comparable
func Or() error {
	type (
		V struct {
			Val int
		}
	)
	var (
		ints = []int{0, 0, 99, 1, 0}
		strs = []string{"", "hello", "world", ""}
		objs = []*V{nil, {999}, nil}
	)

	output.Stdoutl("[non-zero]", cmp.Or(ints...))
	output.Stdoutl("[non-zero]", cmp.Or(strs...))
	output.Stdoutl("[non-zero]", cmp.Or(objs...))

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: cmpop_or

		[Name] "cmpop_or"
		[non-zero]           99
		[non-zero]           hello
		[non-zero]           &{999}


		[Elapsed] 53.06µs
	*/
}
