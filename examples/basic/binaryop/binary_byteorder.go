package binaryop

import (
	"encoding/binary"

	"github.com/devlights/gomy/output"
)

// ByteOrder -- encoding/binary パッケージを用いて Go におけるバイトオーダーの確認をするサンプルです.
func ByteOrder() error {
	// --------------------------------------------------------------
	// Go で バイトオーダー を扱う場合は encoding/binary パッケージを利用する
	//
	// encoding/binary パッケージ
	//   https://golang.org/encoding/binary/
	//
	// バイトオーダーに関しては以下を参照
	//   https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%B3%E3%83%87%E3%82%A3%E3%82%A2%E3%83%B3
	// --------------------------------------------------------------

	// --------------------------------------------------
	// binary.[Big|Little]Endian.PutXXX([]byte, X)
	//   数値を各バイトオーダーで並べたバイト配列にする
	// binary.[Big|Little]Endian.XXX([]byte)
	//   指定したバイト配列を各バイトオーダーで解釈して数値にする
	// --------------------------------------------------
	b32 := make([]byte, 4)
	b16 := make([]byte, 2)
	d32 := 0x0A0B0C0D
	d16 := 0x0A0B

	output.Stdoutl("d32(0x0A0B0C0D)", d32)
	output.Stdoutl("d16(0x0A0B    )", d16)
	output.StdoutHr()

	binary.BigEndian.PutUint32(b32, uint32(d32))
	binary.BigEndian.PutUint16(b16, uint16(d16))

	output.Stdoutl("BigEndian.PutUint32", b32)
	output.Stdoutl("BigEndian.PutUint16", b16)

	output.Stdoutl("BigEndian.Uint32", binary.BigEndian.Uint32(b32))
	output.Stdoutl("BigEndian.Uint16", binary.BigEndian.Uint16(b16))
	output.StdoutHr()

	b32 = make([]byte, 4)
	b16 = make([]byte, 2)

	binary.LittleEndian.PutUint32(b32, uint32(d32))
	binary.LittleEndian.PutUint16(b16, uint16(d16))

	output.Stdoutl("LittleEndian.PutUint32", b32)
	output.Stdoutl("LittleEndian.PutUint16", b16)

	output.Stdoutl("LittleEndian.Uint32", binary.LittleEndian.Uint32(b32))
	output.Stdoutl("LittleEndian.Uint16", binary.LittleEndian.Uint16(b16))
	output.StdoutHr()

	return nil
}
