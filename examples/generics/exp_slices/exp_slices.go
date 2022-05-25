package exp_slices

import (
	"github.com/devlights/gomy/output"
	"golang.org/x/exp/slices"
)

func ExpSlices() error {
	var (
		s1 = []string{"hello", "world"}
		s2 = []string{"hello", "world"}
		s3 = []string{"world", "hello"}
		s4 = []int{100, 101}
	)

	output.Stdoutl("[Equal(s1, s2)]", slices.Equal(s1, s2))
	output.Stdoutl("[Equal(s2, s3)]", slices.Equal(s2, s3))
	// compile error
	//fmt.Println(slices.Equal(s1, s4))

	s5 := slices.Insert(s4, 1, 999)
	output.Stdoutl("[Insert]", s4, s5)

	idx := slices.Index(s5, 999)
	s6 := slices.Delete(s5, idx, idx+1)
	output.Stdoutl("[Delete]", s5, s6)

	return nil
}
