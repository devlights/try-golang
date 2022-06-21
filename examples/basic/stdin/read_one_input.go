package stdin

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// ReadOneInput -- fmt.Scan() を利用して標準入力から一つの入力値を受け取るサンプルです。
//
// REFERENCES:
//   - https://dev.to/itnext/go-from-the-beginning-reading-user-input-i79
func ReadOneInput() error {
	var (
		message        string
		readTokenCount int
		err            error
	)

	// fmt.Scan() は渡した引数を変更するので ポインタ で渡す必要がある
	fmt.Print("ENTER MESSAGE: ")
	readTokenCount, err = fmt.Scan(&message)
	if err != nil {
		return err
	}

	output.Stdoutf("[fmt.Scan]", "read-token-count=%d, message=%s\n", readTokenCount, message)

	return nil
}
