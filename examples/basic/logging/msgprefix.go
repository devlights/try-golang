package logging

import (
	"bytes"
	"log"

	"github.com/devlights/gomy/output"
)

// Msgprefix -- log.Lmsgprefix フラグのサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/log@go1.17.6#pkg-constants
//   - https://tutuz-tech.hatenablog.com/entry/2021/01/30/192956
//   - https://qiita.com/immrshc/items/a080975c6c7e23498944
func Msgprefix() error {
	var (
		buf1 = &bytes.Buffer{}
		buf2 = &bytes.Buffer{}
		log1 = log.New(buf1, "TRACE\t", log.LstdFlags)                // 行の先頭にprefixが出力される
		log2 = log.New(buf2, "DEBUG\t", log.LstdFlags|log.Lmsgprefix) // メッセージの手前にprefixが出力される
	)

	for _, l := range []*log.Logger{log1, log2} {
		l.Println("hello world")
	}

	output.Stdoutf("[-Lmsgprefix]", buf1.String())
	output.Stdoutf("[+Lmsgprefix]", buf2.String())

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: logging_msgprefix

	   [Name] "logging_msgprefix"
	   [-Lmsgprefix]        TRACE      2024/01/22 02:33:14 hello world
	   [+Lmsgprefix]        2024/01/22 02:33:14 DEBUG  hello world


	   [Elapsed] 188.59µs
	*/

}
