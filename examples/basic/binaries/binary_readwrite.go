package binaries

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
	//   https://golang.org/pkg/encoding/binary/
	//
	// バイトオーダーに関しては以下を参照
	//   https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%B3%E3%83%87%E3%82%A3%E3%82%A2%E3%83%B3
	//
	// また、Go1.9より math/bits パッケージも追加された
	// このパッケージは、ビット操作用のパッケージとなっている
	// --------------------------------------------------------------
	output.StdoutHr()
	output.Stdoutl("[BigEndian]")
	output.StdoutHr()

	if err := readBigEndian(); err != nil {
		return err
	}

	output.StdoutHr()
	output.Stdoutl("[LittleEndian]")
	output.StdoutHr()

	if err := readLittleEndian(); err != nil {
		return err
	}

	output.StdoutHr()

	return nil
}

func readBigEndian() error {
	const (
		MaxDataCount = 7
	)

	type (
		header struct {
			MsgNo     uint16
			DataCount byte
			Dummy     byte
		}

		data struct {
			Id    uint32
			Value [4]byte
		}

		msg struct {
			Header header
			Data   [MaxDataCount]data
		}
	)

	b := []byte{
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

	r := bytes.NewReader(b)
	d := msg{}

	if err := binary.Read(r, binary.BigEndian, &d); err != nil {
		return err
	}

	output.Stdoutl("binary.Read(BigEndian)", d)
	for i := 0; i < int(d.Header.DataCount); i++ {
		output.Stdoutf("data", ">>> %v\n", d.Data[i])
	}

	return nil
}

func readLittleEndian() error {
	const (
		MaxDataCount = 7
	)

	type (
		header struct {
			MsgNo     uint16
			DataCount byte
			Dummy     byte
		}

		data struct {
			Id    uint32
			Value [4]byte
		}

		msg struct {
			Header header
			Data   [MaxDataCount]data
		}
	)

	b := []byte{
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

	r := bytes.NewReader(b)
	d := msg{}

	if err := binary.Read(r, binary.LittleEndian, &d); err != nil {
		return err
	}

	output.Stdoutl("binary.Read(LittleEndian)", d)
	for i := 0; i < int(d.Header.DataCount); i++ {
		output.Stdoutf("data", ">>> %v\n", d.Data[i])
	}

	return nil
}
