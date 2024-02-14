package slices

import (
	"github.com/devlights/gomy/output"
)

// ConvertToArrayGo117 は、Go 1.17 以降で有効な スライス から 配列 への変換方法についてのサンプルです。
//
// # REFERENCES
//   - https://tip.golang.org/ref/spec#Conversions_from_slice_to_array_pointer
//   - https://www.jetbrains.com/go/guide/tips/go-1-17-convert-slice-to-array-pointer/
//   - https://zenn.dev/koya_iwamura/articles/bb9b590b57d825
func ConvertToArrayGo117() error {
	//
	// Go 1.17 から以下のやり方でスライスから配列へ変換できるようになった
	// 当然であるが、ポインタで扱うか否かで挙動が変わるので注意
	//

	var (
		original            = []string{"golang", "python", "csharp", "java", "dart", "rust", "javascript"}
		array1   [5]string  = *(*[5]string)(original[:5])
		array2   *[2]string = (*[2]string)(original[5:])
	)
	output.Stdoutf("[array1  ]", "%v\n", array1)
	output.Stdoutf("[array2  ]", "%v\n", array2)
	output.Stdoutf("[original]", "%v\n", original)
	output.StdoutHr()

	array1[0] = "*****"
	array2[0] = "*****"

	output.Stdoutf("[array1  ]", "%v\n", array1)
	output.Stdoutf("[array2  ]", "%v\n", array2)
	output.Stdoutf("[original]", "%v\n", original)

	// 以下は、コンパイルは通るが実行時エラーとなる。
	// 	panic: runtime error: cannot convert slice with length 7 to pointer to array with length 100
	//
	// var (
	// 	array3 = *(*[100]string)(original)
	// )
	// output.Stdoutf("[array3  ]", "%v\n", array3)

	// 以下は、コンパイルは通るが実行時エラーとなる。
	// 	panic: runtime error: cannot convert slice with length 7 to pointer to array with length 100
	// var (
	// 	array4 = (*[100]string)(original)
	// )
	// output.Stdoutf("[array4  ]", "%v\n", array4)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_convert_to_array_go117

	   [Name] "slice_convert_to_array_go117"
	   [array1  ]           [golang python csharp java dart]
	   [array2  ]           &[rust javascript]
	   [original]           [golang python csharp java dart rust javascript]
	   --------------------------------------------------
	   [array1  ]           [***** python csharp java dart]
	   [array2  ]           &[***** javascript]
	   [original]           [golang python csharp java dart ***** javascript]


	   [Elapsed] 83.89µs
	*/

}
