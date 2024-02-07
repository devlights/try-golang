package scannerop

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/devlights/gomy/output"
)

// Custom -- bufio.Scanner に 独自の bufio.SplitFunc を指定して処理するサンプルです.
//
// REFERENCES
//   - https://pkg.go.dev/bufio@go1.18.1#Scanner
//   - https://pkg.go.dev/bufio@go1.18.1#SplitFunc
func Custom() error {
	var (
		gen = func(delimiter []byte) bufio.SplitFunc {
			return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
				if atEOF && len(data) == 0 {
					// 処理するべきデータがない
					return 0, nil, nil
				}

				if i := bytes.Index(data, delimiter); i >= 0 {
					// 見つかったので、Scanner 側にトークンを渡す
					return i + len(delimiter), data[0:i], nil
				}

				if atEOF {
					// 最後のデータを纏めて Scanner 側に渡す
					return len(data), data, nil
				}

				// このターンではトークンが見つからなかったので、更に読み進むように Scanner に通知
				return 0, nil, nil
			}
		}
	)

	for i, fn := range []bufio.SplitFunc{bufio.ScanWords, gen([]byte("[TAB]")), gen([]byte("{TAB}"))} {
		var (
			r = strings.NewReader("hello[TAB]world[TAB]こんにちわ{TAB}世界")
			s = bufio.NewScanner(r)
		)

		s.Split(fn)

		for s.Scan() {
			prefix := fmt.Sprintf("[%d]", i)
			output.Stdoutl(prefix, s.Text())
		}
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: scannerop_custom

	   [Name] "scannerop_custom"
	   [0]                  hello[TAB]world[TAB]こんにちわ{TAB}世界
	   [1]                  hello
	   [1]                  world
	   [1]                  こんにちわ{TAB}世界
	   [2]                  hello[TAB]world[TAB]こんにちわ
	   [2]                  世界


	   [Elapsed] 99µs
	*/

}
