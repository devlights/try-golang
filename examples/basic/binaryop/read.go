package binaryop

import (
	"bytes"
	"encoding/binary"

	"github.com/devlights/gomy/output"
)

// Read は、バイナリを読み込むサンプルです.
func Read() error {
	var (
		bin = []byte{
			0, 0, 0, 1,
			2, 0, 0, 0,

			0x68, 0x65, 0x6c, 0x6c, 0x6f,
			// h    e     l     l     o
			0x77, 0x6f, 0x72, 0x6c, 0x64,
			// w    o     r     l     d

		}
		buf = bytes.NewBuffer(bin)
	)

	var (
		i   int32
		err error
	)

	if err = binary.Read(buf, binary.BigEndian, &i); err != nil {
		return err
	}

	var (
		i2 int32
	)

	if err = binary.Read(buf, binary.LittleEndian, &i2); err != nil {
		return nil
	}

	var (
		hello = make([]byte, 5)
		world = make([]byte, 5)
	)

	if err = binary.Read(buf, binary.BigEndian, &hello); err != nil {
		return err
	}

	if _, err = buf.Read(world); err != nil {
		return err
	}

	output.Stdoutl("[bin   ]", bin)
	output.Stdoutl("[output]", i, i2, string(hello), string(world))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: binaryop_read

	   [Name] "binaryop_read"
	   [bin   ]             [0 0 0 1 2 0 0 0 104 101 108 108 111 119 111 114 108 100]
	   [output]             1 2 hello world


	   [Elapsed] 39.44µs
	*/

}
