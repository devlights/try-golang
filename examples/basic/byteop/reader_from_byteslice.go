package byteop

import (
	"bytes"
	"io"

	"github.com/devlights/gomy/output"
)

// ReaderFromByteSlice -- []byte から io.Reader を生成するサンプルです.
func ReaderFromByteSlice() error {
	// -----------------------------------------------------------------
	// []byte から io.Reader を得るには
	//		bytes.NewReader()
	// を使う.
	// -----------------------------------------------------------------
	var (
		b = []byte{1, 2, 3, 4, 5}
	)

	reader := bytes.NewReader(b)
	for {
		v, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				// 最後までデータを読み込んだ
				break
			}

			return err
		}

		output.Stdoutl("[ReadByte]", v)
	}

	return nil
}
