package rand

import (
	"crypto/rand"

	"github.com/devlights/gomy/iter"
	"github.com/devlights/gomy/output"
)

// Read -- crypto/rand.Read を用いてセキュリティ的に安全な乱数を生成するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/crypto/rand@go1.19#pkg-variables
//   - https://github.com/golang/go/wiki/CodeReviewComments#crypto-rand
func Read() error {

	for range iter.Range(10) {
		var (
			buf = make([]byte, 32)
			err error
		)

		if _, err = rand.Read(buf); err != nil {
			return err
		}

		output.Stdoutf("[Crypto][rand.Read]", "%x\n", buf)
	}

	return nil
}
