package sets

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

func Set03() error {

	// Set.Difference の 確認 (差集合)
	s1 := mapset.NewSet("hello", "world")
	s2 := mapset.NewSet("golang", "world", "python")

	s3 := s1.Difference(s2)

	fmt.Println(s3)

	return nil
}
