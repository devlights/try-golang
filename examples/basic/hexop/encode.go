package hexop

import (
	"encoding/hex"

	"github.com/devlights/gomy/output"
)

// Encode -- encoding/hex.Encode を利用したサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/encoding/hex@go1.17.6
func Encode() error {
	// Go では 何かのデータ を変換する際は基本 []byte でやり取りする
	// エンコードとは、元のデータを16進数に変換すること

	var (
		src    = []byte("hello world")
		encLen = hex.EncodedLen(len(src))
		dst    = make([]byte, encLen)
	)

	hex.Encode(dst, src)

	output.Stdoutl("original", string(src))
	output.Stdoutl("enc-len ", encLen)
	output.Stdoutl("encoded ", string(dst))

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: hexop_encode

	   [Name] "hexop_encode"
	   original             hello world
	   enc-len              22
	   encoded              68656c6c6f20776f726c64


	   [Elapsed] 23.54µs
	*/

}
