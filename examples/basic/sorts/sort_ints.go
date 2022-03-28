package sorts

import (
	"sort"

	"github.com/devlights/gomy/output"
)

// Ints -- sort.Ints() のサンプルです.
//
// REFERENCES
//   - https://dev.to/jpoly1219/easy-sorting-in-go-56ae
//   - https://pkg.go.dev/sort
func Ints() error {
	var (
		original = []int{5, 2, 3, 4, 1}
		sorted   = make([]int, len(original))
	)

	copy(sorted, original)
	sort.Ints(sorted)

	output.Stdoutl("[original]", original)
	output.Stdoutl("[sorted  ]", sorted)

	return nil
}
