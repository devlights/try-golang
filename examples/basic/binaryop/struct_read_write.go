package binaryop

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"

	"github.com/devlights/gomy/output"
)

// StructReadWrite は、構造体をバイナリにパックし、それをバイナリとして書き出すサンプルです.
//
// # REFERENCES
//   - https://devlights.hatenablog.com/entry/2021/06/16/014722
//   - https://pkg.go.dev/encoding/binary@go1.20.2#Read
//   - https://pkg.go.dev/encoding/binary@go1.20.2#Write
func StructReadWrite() error {
	type St struct {
		Num  uint32
		S    [10]byte
		Num2 uint16
	}

	//
	// 以下のコードでは、エラー処理は省略している
	//

	// バイナリを生成
	var (
		buf = bytes.NewBuffer(nil)
	)

	binary.Write(buf, binary.BigEndian, uint32(0xDEADBEEF))
	binary.Write(buf, binary.BigEndian, []byte("HELLOWORLD"))
	binary.Write(buf, binary.BigEndian, uint16(0xCAFE))

	// バイナリを読み込み、構造体にパック
	var (
		bin = buf.Bytes()
		st  St
	)

	binary.Read(buf, binary.BigEndian, &st)

	// 構造体の値をバイナリとして出力
	buf.Reset()
	binary.Write(buf, binary.BigEndian, &st)

	// HEXダンプして確認
	var (
		dumpBin = func(b []byte) {
			w := hex.Dumper(output.Writer())
			defer w.Close()

			w.Write(b)
		}
		dumpSt = func(st *St) {
			w := hex.Dumper(output.Writer())
			defer w.Close()

			binary.Write(w, binary.BigEndian, st.Num)
			binary.Write(w, binary.BigEndian, st.S)
			binary.Write(w, binary.BigEndian, st.Num2)
		}
	)

	dumpBin(bin)
	output.StdoutHr()
	dumpSt(&st)
	output.StderrHr()
	dumpBin(buf.Bytes())

	return nil
}
