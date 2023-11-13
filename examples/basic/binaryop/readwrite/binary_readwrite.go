package readwrite

import (
	"bytes"
	"encoding/binary"

	"github.com/devlights/gomy/output"
)

// ReadWrite -- binary パッケージの Read/Write を利用してエンディアン指定でデータを読み込むサンプルです.
func ReadWrite() error {
	// --------------------------------------------------------------
	// Go で バイトオーダー を扱う場合は encoding/binary パッケージを利用する
	//
	// encoding/binary パッケージ
	//   https://golang.org/encoding/binary/
	//
	// バイトオーダーに関しては以下を参照
	//   https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%B3%E3%83%87%E3%82%A3%E3%82%A2%E3%83%B3
	//
	// また、Go1.9より math/bits パッケージも追加された
	// このパッケージは、ビット操作用のパッケージとなっている
	// --------------------------------------------------------------
	var (
		bigEndianBytes = []byte{
			// header {MsgNo=10,DataCount=4}
			0x00, 0x0A, 0x04, 0x00,
			// data {
			//   (Id=11,Value={1,2,3,4})
			//   (Id=12,Value={5,6,7,8})
			//   (Id=13,Value={9,10,11,12})
			//   (Id=16909060,Value={13,14,15,16})
			//   (Id=0,Value={0,0,0,0})
			//   (Id=0,Value={0,0,0,0})
			//   (Id=0,Value={0,0,0,0})
			// }
			//
			// 0x01020304 = 16909060
			0x00, 0x00, 0x00, 0x0B, 0x01, 0x02, 0x03, 0x04,
			0x00, 0x00, 0x00, 0x0C, 0x05, 0x06, 0x07, 0x08,
			0x00, 0x00, 0x00, 0x0D, 0x09, 0x0A, 0x0B, 0x0C,
			0x01, 0x02, 0x03, 0x04, 0x0D, 0x0E, 0x0F, 0x10,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}

		littleEndianBytes = []byte{
			// header {MsgNo=10,DataCount=4}
			0x0A, 0x00, 0x04, 0x00,
			// data {
			//   (Id=11,Value={1,2,3,4})
			//   (Id=12,Value={5,6,7,8})
			//   (Id=13,Value={9,10,11,12})
			//   (Id=16909060,Value={13,14,15,16})
			//   (Id=0,Value={0,0,0,0})
			//   (Id=0,Value={0,0,0,0})
			//   (Id=0,Value={0,0,0,0})
			// }
			//
			// 0x01020304 = 16909060
			0x0B, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04,
			0x0C, 0x00, 0x00, 0x00, 0x05, 0x06, 0x07, 0x08,
			0x0D, 0x00, 0x00, 0x00, 0x09, 0x0A, 0x0B, 0x0C,
			0x04, 0x03, 0x02, 0x01, 0x0D, 0x0E, 0x0F, 0x10,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}
	)

	output.StdoutHr()
	output.Stdoutl("[BigEndian] Read")
	output.StdoutHr()

	msg, err := readBigEndian(bigEndianBytes)
	if err != nil {
		return err
	}

	output.StdoutHr()
	output.Stdoutl("[BigEndian] Write")
	output.StdoutHr()

	if err = writeBigEndian(msg, bigEndianBytes); err != nil {
		return err
	}

	output.StdoutHr()
	output.Stdoutl("[LittleEndian] Read")
	output.StdoutHr()

	msg, err = readLittleEndian(littleEndianBytes)
	if err != nil {
		return err
	}

	output.StdoutHr()
	output.Stdoutl("[LittleEndian] Write")
	output.StdoutHr()

	if err = writeLittleEndian(msg, littleEndianBytes); err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: binaryop_readwrite

	   [Name] "binaryop_readwrite"
	   --------------------------------------------------
	   [BigEndian] Read
	   --------------------------------------------------
	   binary.Read(BigEndian) {{10 4 0} [{11 [1 2 3 4]} {12 [5 6 7 8]} {13 [9 10 11 12]} {16909060 [13 14 15 16]} {0 [0 0 0 0]} {0 [0 0 0 0]} {0 [0 0 0 0]}]}
	   data                 >>> {11 [1 2 3 4]}
	   data                 >>> {12 [5 6 7 8]}
	   data                 >>> {13 [9 10 11 12]}
	   data                 >>> {16909060 [13 14 15 16]}
	   --------------------------------------------------
	   [BigEndian] Write
	   --------------------------------------------------
	   [binary.write]       [0 10 4 0 0 0 0 11 1 2 3 4 0 0 0 12 5 6 7 8 0 0 0 13 9 10 11 12 1 2 3 4 13 14 15 16 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	   [orignal     ]       [0 10 4 0 0 0 0 11 1 2 3 4 0 0 0 12 5 6 7 8 0 0 0 13 9 10 11 12 1 2 3 4 13 14 15 16 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	   --------------------------------------------------
	   [LittleEndian] Read
	   --------------------------------------------------
	   binary.Read(LittleEndian) {{10 4 0} [{11 [1 2 3 4]} {12 [5 6 7 8]} {13 [9 10 11 12]} {16909060 [13 14 15 16]} {0 [0 0 0 0]} {0 [0 0 0 0]} {0 [0 0 0 0]}]}
	   data                 >>> {11 [1 2 3 4]}
	   data                 >>> {12 [5 6 7 8]}
	   data                 >>> {13 [9 10 11 12]}
	   data                 >>> {16909060 [13 14 15 16]}
	   --------------------------------------------------
	   [LittleEndian] Write
	   --------------------------------------------------
	   [binary.write]       [10 0 4 0 11 0 0 0 1 2 3 4 12 0 0 0 5 6 7 8 13 0 0 0 9 10 11 12 4 3 2 1 13 14 15 16 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	   [orignal     ]       [10 0 4 0 11 0 0 0 1 2 3 4 12 0 0 0 5 6 7 8 13 0 0 0 9 10 11 12 4 3 2 1 13 14 15 16 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]


	   [Elapsed] 495.811µs
	*/

}

func writeBigEndian(msg *message, b []byte) error {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, msg); err != nil {
		return err
	}

	output.Stdoutl("[binary.write]", buf.Bytes())
	output.Stdoutl("[orignal     ]", b)

	return nil
}

func writeLittleEndian(msg *message, b []byte) error {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, msg); err != nil {
		return err
	}

	output.Stdoutl("[binary.write]", buf.Bytes())
	output.Stdoutl("[orignal     ]", b)

	return nil
}

func readBigEndian(b []byte) (*message, error) {
	r := bytes.NewReader(b)
	d := message{}

	if err := binary.Read(r, binary.BigEndian, &d); err != nil {
		return &d, err
	}

	output.Stdoutl("binary.Read(BigEndian)", d)
	for i := 0; i < int(d.Header.DataCount); i++ {
		output.Stdoutf("data", ">>> %v\n", d.Data[i])
	}

	return &d, nil
}

func readLittleEndian(b []byte) (*message, error) {
	r := bytes.NewReader(b)
	d := message{}

	if err := binary.Read(r, binary.LittleEndian, &d); err != nil {
		return nil, err
	}

	output.Stdoutl("binary.Read(LittleEndian)", d)
	for i := 0; i < int(d.Header.DataCount); i++ {
		output.Stdoutf("data", ">>> %v\n", d.Data[i])
	}

	return &d, nil
}
