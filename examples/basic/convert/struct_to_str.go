package convert

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

type (
	sts struct {
		id   int
		name string
	}

	yen int
)

func (s sts) String() string {
	return fmt.Sprintf("id=%d,name=%s", s.id, s.name)
}

func (y yen) String() string {
	return fmt.Sprintf("¥%d", y)
}

// StructToStr は、fmt.Sprint() を利用して 構造体 を 文字列 にするサンプルです.
func StructToStr() error {
	var (
		s1 = sts{
			id:   100,
			name: "sts",
		}
		y = yen(1000)
	)

	var (
		fn = func(o fmt.Stringer) {
			output.Stdoutf("[struct to str]", "%[1]v(%[1]T) --> %[2]q(%[2]T)\n", o, fmt.Sprint(o))
		}
	)

	fn(s1)
	fn(y)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: convert_struct_to_str

	   [Name] "convert_struct_to_str"
	   [struct to str]      id=100,name=sts(convert.sts) --> "id=100,name=sts"(string)
	   [struct to str]      ¥1000(convert.yen) --> "¥1000"(string)


	   [Elapsed] 55.5µs
	*/

}
