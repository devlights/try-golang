package scanop

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// ReadOneInput は、fmt.Scan() で一つの値を読み取るサンプルです.
//
// # REFERENCES
//
//   - https://dev.to/azure/go-from-the-beginning-reading-user-input-i79
//   - https://pkg.go.dev/fmt@go1.19.2#Scan
func ReadOneInput() error {
	var (
		value string
	)

	fmt.Print("INPUT: ")

	// スペースで区切られて読み取られるため、例えば"hello world"と指定するとhelloだけが読み取られる.
	n, err := fmt.Scan(&value)
	if err != nil {
		return err
	}

	output.Stdoutf("[fmt.Scan]", "count=%d\tvalue=%v\n", n, value)

	return nil
}
