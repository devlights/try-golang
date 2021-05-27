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
}
