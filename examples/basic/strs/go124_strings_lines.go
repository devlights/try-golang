package strs

import (
	"fmt"
	"iter"
	"os"
	"strings"
	"unsafe"
)

// Go124StringsLines は、Go 1.24 で追加された strings.Lines() のサンプルです.
//
// Go 1.23 で追加されたイテレータを返すようになっています。
// strings.Lines()は、文字列を '\n' で区切ったデータをイテレータで提供してくれます。
// このデータには '\n' 自身も付いた状態で取得出来るので、その点には注意が必要。
// (fmt.Println()でそのまま出力すると２個改行が入ることになる)
//
// # REFERENCES
//   - https://pkg.go.dev/strings@go1.24.0#Lines
//   - https://tip.golang.org/doc/go1.24#stringspkgstrings
//   - https://tip.golang.org/doc/go1.23#iterators
func Go124StringsLines() error {
	var (
		// 最初のN行のみに絞るためのイテレータ関数
		firstN = func(n int, lines iter.Seq[string]) iter.Seq[string] {
			return func(yield func(string) bool) {
				count := 0
				for line := range lines {
					if count >= n || !yield(line) {
						return
					}

					count++
				}
			}
		}
		// 前に番号を付与するイテレータ
		withLineNum = func(lines iter.Seq[string]) iter.Seq[string] {
			return func(yield func(string) bool) {
				count := 1
				for line := range lines {
					if !yield(fmt.Sprintf("%2d: %s", count, line)) {
						return
					}

					count++
				}
			}
		}
		// []byteからstringへ変換するためのヘルパー関数
		toStr = func(b []byte) string {
			return unsafe.String(unsafe.SliceData(b), len(b))
		}

		data, _ = os.ReadFile("README.md")   // ファイルデータを読み出して
		lines   = strings.Lines(toStr(data)) // 行単位にして
		first5  = firstN(5, lines)           // 先頭５行のみにして
		withNum = withLineNum(first5)        // 行番号を付与
	)

	for line := range withNum {
		fmt.Printf("%s", line)
	}

	return nil

	/*
	   $ task
	   task: [build] go build -o "/workspace/try-golang/try-golang" .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: string_go124_strings_lines

	   [Name] "string_go124_strings_lines"

	   	1:
	   	2: # try-golang
	   	3:
	   	4: This is my TUTORIAL project for golang.
	   	5:

	   [Elapsed] 94.12µs
	*/
}
