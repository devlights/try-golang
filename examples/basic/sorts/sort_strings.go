package sorts

import (
	"sort"

	"github.com/devlights/gomy/output"
)

// Strings -- sort.Strings() のサンプルです.
//
// REFERENCES
//   - https://dev.to/jpoly1219/easy-sorting-in-go-56ae
//   - https://pkg.go.dev/sort
func Strings() error {
	var (
		original = []string{"z", "e", "u", "s"}
		sorted   = make([]string, len(original))
	)

	copy(sorted, original)
	sort.Strings(sorted)

	output.Stdoutl("[original]", original)
	output.Stdoutl("[sorted  ]", sorted)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: sort_strings

	   [Name] "sort_strings"
	   [original]           [z e u s]
	   [sorted  ]           [e s u z]


	   [Elapsed] 25.99µs
	*/

}
