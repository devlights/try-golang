package sets

import (
	"fmt"
	"github.com/deckarep/golang-set"
)

func Set02() error {

	// Set.Union の 確認
	s1 := mapset.NewSet("hello", "world")
	s2 := mapset.NewSet("golang", "world", "python")

	s3 := s1.Union(s2)

	fmt.Println(s3)

	return nil
}
