package hexop

import (
	"bytes"
	"encoding/hex"

	"github.com/devlights/gomy/output"
)

func Decoder() error {
	var (
		src = "68656c6c6f20776f726c64"
		buf = bytes.NewBufferString(src)
		dec = hex.NewDecoder(buf)
		dst = make([]byte, hex.DecodedLen(buf.Len()))
	)

	if _, err := dec.Read(dst); err != nil {
		return err
	}

	output.Stdoutl("original", src)
	output.Stdoutl("decoded ", string(dst))

	return nil
}
