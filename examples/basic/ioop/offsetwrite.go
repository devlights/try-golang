package ioop

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	"github.com/devlights/gomy/errs"
	"github.com/devlights/gomy/output"
)

// OffsetWrite は、io.OffsetWriter を利用したサンプルです。
//
// io.OffsetWriter は、io.WriterAt インターフェースを実装しているものを要求する。
//
// os.File が io.WriterAt を実装している。
//
// # REFERENCES
//   - https://pkg.go.dev/io@go1.22.2#OffsetWriter
//   - https://cs.opensource.google/go/go/+/go1.22.2:src/io/io.go;l=570
func OffsetWrite() error {
	w, err := os.CreateTemp(os.TempDir(), "trygolang")
	if err != nil {
		return err
	}
	defer os.Remove(w.Name())

	output.Stdoutl("[File]", w.Name())

	fstat := errs.Drop(w.Stat())
	{
		defer w.Close()

		bufW := bufio.NewWriter(w)
		bufW.WriteString("helloworld こんにちは世界")
		bufW.Flush()

		offW := io.NewOffsetWriter(w, 11) // "helloworld " の次の位置（つまり「こ」）にオフセットを設定
		offW.Write([]byte("コンニチハ"))

		offW.Seek(int64(15), io.SeekStart) // "コンニチハ" の次の位置（つまり「世」）にシークポジションをセット
		offW.Write([]byte("セカイ"))
	}

	data := errs.Drop(os.ReadFile(filepath.Join(os.TempDir(), fstat.Name())))
	output.Stdoutl("[offW]", string(data))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: ioop_offset_write

	   [Name] "ioop_offset_write"
	   [offW]               helloworld コンニチハセカイ


	   [Elapsed] 165.78µs
	*/

}
