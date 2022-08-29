package rand

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/devlights/gomy/iter"
	"github.com/devlights/gomy/output"
)

// RandomBytes -- crypto/rand.Reader を用いてセキュリティ的に安全な乱数を生成するサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/crypto/rand@go1.19#pkg-variables
func RandomBytes() error {
	var (
		r = rand.Reader
	)

	for range iter.Range(10) {
		var (
			b   = make([]byte, 32)
			err error
		)

		_, err = io.ReadFull(r, b)
		if err != nil {
			return err
		}

		output.Stdoutl("[Crypto][rand.Reader]", hex.EncodeToString(b))
	}

	return nil
}
