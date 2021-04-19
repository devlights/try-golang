package convert

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

type sts struct {
	id   int
	name string
}

func (s sts) String() string {
	return fmt.Sprintf("id=%d,name=%s", s.id, s.name)
}

// StructToStr は、fmt.Sprint() を利用して 構造体 を 文字列 にするサンプルです.
func StructToStr() error {

	s1 := sts{
		id:   100,
		name: "sts",
	}

	s := fmt.Sprint(s1)
	output.Stdoutf("[struct to str]", "%[1]v(%[1]T) --> %[2]q(%[2]T)\n", s1, s)

	return nil
}
