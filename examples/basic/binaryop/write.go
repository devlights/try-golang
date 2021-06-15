package binaryop

import (
	"bytes"
	"encoding/binary"

	"github.com/devlights/gomy/output"
)

// Write は、バイナリを書き込むサンプルです.
func Write() error {
	var (
		buf = new(bytes.Buffer)
		err error
	)

	if err = binary.Write(buf, binary.BigEndian, int32(1)); err != nil {
		return err
	}

	if err = binary.Write(buf, binary.LittleEndian, int32(2)); err != nil {
		return err
	}

	if err = binary.Write(buf, binary.BigEndian, []byte("helloworld")); err != nil {
		return err
	}

	output.Stdoutl("[output]", buf.Bytes())

	return nil
}
