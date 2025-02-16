package byteop

import (
	"bytes"
	"fmt"
	"iter"
	"os"
)

// Go124BytesLines は、Go 1.24 で追加された bytes.Lines() のサンプルです.
//
// Go 1.23 で追加されたイテレータを返すようになっています。
// bytes.Lines()は、バイト列を '\n' で区切ったデータをイテレータで提供してくれます。
// このデータには '\n' 自身も付いた状態で取得出来るので、その点には注意が必要。
// (fmt.Println()でそのまま出力すると２個改行が入ることになる)
//
// # REFERENCES
//   - https://pkg.go.dev/bytes@go1.24.0#Lines
//   - https://tip.golang.org/doc/go1.24#bytespkgbytes
//   - https://tip.golang.org/doc/go1.23#iterators
func Go124BytesLines() error {
	var (
		// 最初のN行のみに絞るためのイテレータ関数
		firstN = func(n int, lines iter.Seq[[]byte]) iter.Seq[[]byte] {
			return func(yield func([]byte) bool) {
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
		withLineNum = func(lines iter.Seq[[]byte]) iter.Seq[string] {
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

		data, _ = os.ReadFile("README.md") // ファイルデータを読み出して
		lines   = bytes.Lines(data)        // 行単位にして
		first5  = firstN(5, lines)         // 先頭５行のみにして
		withNum = withLineNum(first5)      // 行番号を付与
	)

	for line := range withNum {
		fmt.Printf("%s", line)
	}

	return nil
}
