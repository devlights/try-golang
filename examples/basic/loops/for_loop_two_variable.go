package loops

import "github.com/devlights/gomy/output"

// ForLoopTwoVariable -- for ループで ２つの変数 を初期化してループさせるサンプルです.
//
// REFERENCES
//   - https://stackoverflow.com/questions/38081807/for-loop-of-two-variables-in-go
func ForLoopTwoVariable() error {
	var (
		s = "helloworld"
		r = make([]byte, len(s))
	)

	// for i:=len(s)-1, j:=0; i >= 0; i--,j++ ではないことに注意
	for i, j := len(s)-1, 0; i >= 0; i, j = i-1, j+1 {
		r[j] = s[i]
	}

	output.Stdoutl("[orig]", s)
	output.Stdoutl("[ret ]", string(r))

	return nil
}
