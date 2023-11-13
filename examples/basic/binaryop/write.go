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

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: binaryop_write

	   [Name] "binaryop_write"
	   [output]             [0 0 0 1 2 0 0 0 104 101 108 108 111 119 111 114 108 100]


	   [Elapsed] 15.53µs
	*/

}
