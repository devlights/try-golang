package ioop

import (
	"bytes"
	"io"
	"strings"

	"github.com/devlights/gomy/output"
)

// SectionRead は、io.SectionReader を利用したサンプルです。
//
// io.SectionReader は、指定した範囲のデータを読み込んでくれる特殊ストリーム。
//
// > SectionReader implements Read, Seek, and ReadAt on a section of an underlying ReaderAt.
//
// > SectionReaderは、Read、Seek、ReadAtを実装しています。
//
// # REFERENCES
//   - https://pkg.go.dev/io@go1.22.2#SectionReader
//   - https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/io/io.go;l=501
func SectionRead() error {
	var (
		r    = strings.NewReader("helloworld こんにちは世界")
		secR = io.NewSectionReader(r, 11, 15)
		buf  = new(bytes.Buffer)
		err  error
	)

	_, err = io.Copy(buf, secR)
	if err != nil {
		return err
	}

	output.Stdoutl("[secR]", buf.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: ioop_section_read

	   [Name] "ioop_section_read"
	   [secR]               こんにちは


	   [Elapsed] 24.97µs
	*/

}
