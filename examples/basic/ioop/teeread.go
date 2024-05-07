package ioop

import (
	"io"
	"os"

	"github.com/devlights/gomy/output"
)

// byteCounter は、書き込まれたバイト数を数える特殊 io.Writer です。
type byteCounter int

func (me *byteCounter) Write(p []byte) (n int, err error) {
	*me = byteCounter(int(*me) + len(p))
	return len(p), nil
}

// TeeRead は、io.TeeReader を利用したサンプルです。
//
// io.TeeReader は、teeコマンドと同じような動きをする。読み取ったデータは戻り値のReaderを経由して取得し、さらに引数で指定したio.Writerにも書き込まれる。
//
// > TeeReader returns a Reader that writes to w what it reads from r. All reads from r performed through it are matched with corresponding writes to w.
//
// > TeeReaderは、rから読み取ったものをwに書き込むReaderを返します。TeeReaderを介して実行されるrからのすべての読み取りは、対応するwへの書き込みと一致します。
//
// io.TeeReaderには、内部バッファを持っていない。
// なので、rから読み取ったデータは即時wに書き込まないとブロックされてしまう。
//
// # REFERENCES
//   - https://pkg.go.dev/io@go1.22.2#TeeReader
//   - https://cs.opensource.google/go/go/+/refs/tags/go1.22.2:src/io/io.go;l=618
func TeeRead() error {
	r, err := os.Open("main.go")
	if err != nil {
		return err
	}
	defer r.Close()

	w := byteCounter(0)

	b, err := io.ReadAll(io.TeeReader(r, &w))
	if err != nil {
		return err
	}

	output.Stdoutl("[r]", string(b))
	output.StdoutHr()
	output.Stdoutl("[w]", int(w))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: ioop_tee_read

	   [Name] "ioop_tee_read"
	   [r]                  package main

	   import "github.com/devlights/try-golang/cmd"

	   func main() {
	           cmd.Execute()
	   }

	   --------------------------------------------------
	   [w]                  91


	   [Elapsed] 131.071µs
	*/

}
