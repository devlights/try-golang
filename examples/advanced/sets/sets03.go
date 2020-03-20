package sets

import (
	"fmt"

	"github.com/deckarep/golang-set"
)

// Set03 -- Set.Difference() の動作確認のサンプルです。
func Set03() error {

	// Set.Difference の 確認 (差集合)
	s1 := mapset.NewSet("hello", "world")
	s2 := mapset.NewSet("golang", "world", "python")

	s3 := s1.Difference(s2)

	fmt.Println(s3) // -> "hello"

	return nil
}
