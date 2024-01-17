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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: hexop_encoder

	   [Name] "hexop_encoder"
	   original             hello world
	   encoded              68656c6c6f20776f726c64


	   [Elapsed] 17.57µs
	*/

}
