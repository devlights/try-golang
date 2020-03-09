package sets

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

func Set04() error {

	// Set.Intersect の 確認 (積集合)
	s1 := mapset.NewSet("hello", "world")
	s2 := mapset.NewSet("golang", "world", "python")

	s3 := s1.Intersect(s2)

	fmt.Println(s3) // -> "world"

	return nil
}
