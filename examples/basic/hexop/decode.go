package hexop

import (
	"encoding/hex"

	"github.com/devlights/gomy/output"
)

// Decode -- encoding/hex.Decode を利用したサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/encoding/hex@go1.17.6
func Decode() error {
	// Go では 何かのデータ を変換する際は基本 []byte でやり取りする
	// デコードとは、16進数データを元のデータに戻すこと

	var (
		src    = []byte("68656c6c6f20776f726c64")
		decLen = hex.DecodedLen(len(src))
		dst    = make([]byte, decLen)
	)

	hex.Decode(dst, src)

	output.Stdoutl("original", string(src))
	output.Stdoutl("dec-len ", decLen)
	output.Stdoutl("decoded ", string(dst))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: hexop_decode

	   [Name] "hexop_decode"
	   original             68656c6c6f20776f726c64
	   dec-len              11
	   decoded              hello world


	   [Elapsed] 35.809µs
	*/

}
