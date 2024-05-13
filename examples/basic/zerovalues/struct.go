package zerovalues

import "github.com/devlights/gomy/output"

// Struct は、Goにおける 構造体 のゼロ値についてのサンプルです.
//
// # REFERENCES
//   - https://go.dev/tour/basics/12
//   - https://brain2life.hashnode.dev/default-zero-values-in-go
func Struct() error {
	//
	// 構造体 の ゼロ値 は、フィールドの型毎にゼロ値が設定された状態.
	//
	type (
		_st struct {
			i int
			b bool
		}
	)

	var (
		st _st
	)

	output.Stdoutf("[struct zerovalue]", "%+v\n", st)

	st.i = 100
	st.b = true
	output.Stdoutf("[struct assign values]", "%+v\n", st)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: zerovalues_struct

	   [Name] "zerovalues_struct"
	   [struct zerovalue]   {i:0 b:false}
	   [struct assign values] {i:100 b:true}


	   [Elapsed] 12.47µs
	*/

}
