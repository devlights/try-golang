package binaryop

import (
	"bytes"
	"encoding/binary"

	"github.com/devlights/gomy/output"
)

// MapStruct は、構造体にバイナリデータをマッピングするサンプルです.
func MapStruct() error {
	type (
		binSt struct {
			I1 uint32
			I2 uint16
			I3 uint8
			I4 uint8
			S  [10]byte
		}
	)

	var (
		bin = []byte{
			0, 0, 0, 1,
			0, 2,
			3,
			0xFF,

			0x68, 0x65, 0x6c, 0x6c, 0x6f,
			// h    e     l     l     o
			0x77, 0x6f, 0x72, 0x6c, 0x64,
			// w    o     r     l     d

		}
		buf = bytes.NewBuffer(bin)
	)

	var (
		st  binSt
		err error
	)

	if err = binary.Read(buf, binary.BigEndian, &st); err != nil {
		return err
	}

	output.Stdoutl("[output]", st)
	output.Stdoutl("[output]", st.I1, st.I2, st.I3, st.I4, string(st.S[:]))

	return nil
}
