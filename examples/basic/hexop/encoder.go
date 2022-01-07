package hexop

import (
	"bytes"
	"encoding/hex"

	"github.com/devlights/gomy/output"
)

// Encoder -- encoding/hex.NewEncoder を利用したサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/encoding/hex@go1.17.6
func Encoder() error {
	var (
		buf = &bytes.Buffer{}
		enc = hex.NewEncoder(buf)
		src = []byte("hello world")
	)

	if _, err := enc.Write(src); err != nil {
		return err
	}

	output.Stdoutl("original", string(src))
	output.Stdoutl("encoded ", buf.String())

	return nil
}
