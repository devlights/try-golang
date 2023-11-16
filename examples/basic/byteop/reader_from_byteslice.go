package byteop

import (
	"bytes"
	"io"
	"os"

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
		b = []byte{49, 50, 51, 52, 53} // ascii コードで 1, 2, 3, 4, 5
	)

	// bytes.NewReader() は、 *bytes.Reader を返す
	// bytes.Reader は、以下のインターフェースを実装している
	//
	//   - io.Reader
	//   - io.ReaderAt
	//   - io.ByteReader
	//   - io.RuneReader
	//   - io.Seeker
	//   - io.ByteScanner
	//   - io.RuneScanner
	//   - io.WriterTo
	//
	// io.Seeker を実装しているので、シーク可能となっている。
	reader := bytes.NewReader(b)

	// io.Reader として利用
	if err := readAsReader(reader); err != nil {
		return err
	}

	if err := backToStart(reader); err != nil {
		return err
	}

	// io.ByteReader として利用
	if err := readAsByteReader(reader); err != nil {
		return err
	}

	if err := backToStart(reader); err != nil {
		return err
	}

	// io.ReaderAt として利用
	if err := readAtReaderAt(reader); err != nil {
		return err
	}

	if err := backToStart(reader); err != nil {
		return err
	}

	// io.WriterTo として利用
	//   バイナリ値を標準出力にそのまま流し込むので
	//   結果は 12345 と表示される
	_, err := reader.WriteTo(os.Stdin)
	if err != nil {
		return err
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: byteop_reader_from_byteslice

	   [Name] "byteop_reader_from_byteslice"
	   [io.Reader]          49
	   [io.Reader]          50
	   [io.Reader]          51
	   [io.Reader]          52
	   [io.Reader]          53
	   [io.ByteReader]      49
	   [io.ByteReader]      50
	   [io.ByteReader]      51
	   [io.ByteReader]      52
	   [io.ByteReader]      53
	   [io.ReaderAt]        [51 52 53]
	   12345

	   [Elapsed] 133.07µs
	*/

}

func backToStart(s io.Seeker) error {
	_, err := s.Seek(0, io.SeekStart)
	return err
}

func readAsReader(r io.Reader) error {
	buf := make([]byte, 1)
	for {
		readBytes, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				// 最後までデータを読み込んだ
				break
			}

			return err
		}

		if readBytes == 0 {
			break
		}

		output.Stdoutl("[io.Reader]", buf[0])
	}

	return nil
}

func readAsByteReader(r io.ByteReader) error {
	for {
		v, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		output.Stdoutl("[io.ByteReader]", v)
	}

	return nil
}

func readAtReaderAt(r io.ReaderAt) error {
	buf := make([]byte, 3)

	_, err := r.ReadAt(buf, 2)
	if err != nil {
		return err
	}

	output.Stdoutl("[io.ReaderAt]", buf)

	return nil
}
