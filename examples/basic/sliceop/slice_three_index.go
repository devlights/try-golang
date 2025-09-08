package sliceop

import "github.com/devlights/gomy/output"

// ThreeIndex は、スライスにて３つのインデックス指定をした場合のサンプルです.
//
// Three-index slices の仕様は、 Go 1.2 にて導入されたもの。
// 3つ目のインデックス指定は、capacity の量を意図的に調整するためにある。
//
// REFERENCES
//   - https://stackoverflow.com/questions/27938177/golang-slice-slicing-a-slice-with-sliceabc
//   - https://stackoverflow.com/questions/12768744/re-slicing-slices-in-golang/18911267#18911267
//   - https://tip.golang.org/doc/go1.2#three_index
//   - https://go.dev/ref/spec#Slice_expressions
func ThreeIndex() error {
	var (
		s = []string{
			"golang",
			"dotnet",
			"java",
			"python",
			"ruby",
		}
		s2 = s[1:3]
		s3 = s[1:3:4]
	)

	output.Stdoutf("[s      ]", "slice=%v\tlen=%v\tcap=%v\n", s, len(s), cap(s))
	output.StdoutHr()

	output.Stdoutl("[len(s2)]", "len => high(2nd) - low(1st) => 3      - 1 => 2")
	output.Stdoutl("[cap(s2)]", "cap => max(3rd)  - low(1st) => len(s) - 1 => 5 - 1 => 4")
	output.Stdoutf("[s2     ]", "slice=%v\tlen=%v\tcap=%v\n", s2, len(s2), cap(s2))

	output.StdoutHr()

	output.Stdoutl("[len(s3)]", "len => high(2nd) - low(1st) => 3 - 1 => 2")
	output.Stdoutl("[cap(s3)]", "cap => max(3rd)  - low(1st) => 4 - 1 => 3")
	output.Stdoutf("[s3     ]", "slice=%v\tlen=%v\tcap=%v\n", s3, len(s3), cap(s3))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_three_index

	   [Name] "slice_three_index"
	   [s      ]            slice=[golang dotnet java python ruby]     len=5   cap=5
	   --------------------------------------------------
	   [len(s2)]            len => high(2nd) - low(1st) => 3      - 1 => 2
	   [cap(s2)]            cap => max(3rd)  - low(1st) => len(s) - 1 => 5 - 1 => 4
	   [s2     ]            slice=[dotnet java]        len=2   cap=4
	   --------------------------------------------------
	   [len(s3)]            len => high(2nd) - low(1st) => 3 - 1 => 2
	   [cap(s3)]            cap => max(3rd)  - low(1st) => 4 - 1 => 3
	   [s3     ]            slice=[dotnet java]        len=2   cap=3


	   [Elapsed] 72.2µs
	*/

}
