package scannerop

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// BufferSize は、bufio.Scanner.Buffer()にてバッファサイズを設定することによる挙動の違いについてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/bufio@go1.25.0#Scanner.Buffer
func BufferSize() error {

	// Scannerのデフォルトバッファサイズ (64KB) を超える文字列を生成
	var (
		s   = strings.Repeat("a", 65536)
		m   = map[string]string{"field": s}
		buf = bytes.NewBuffer(nil)
		enc = json.NewEncoder(buf)
		err error
	)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "")
	if err = enc.Encode(m); err != nil {
		return err
	}

	fmt.Printf("一行のバイトサイズ: %dbytes\n\n", buf.Len())

	// デフォルトバッファ（64KB）で読み込み
	fmt.Println("[デフォルトバッファサイズで読み込み]")
	{
		var (
			scanner = bufio.NewScanner(buf)
		)
		for scanner.Scan() {
			fmt.Fprintln(io.Discard, scanner.Text())
		}

		fmt.Printf("\t結果： %v\n", scanner.Err())
	}

	// バッファを拡張（1MB）して読み込み
	fmt.Println("[拡張バッファサイズで読み込み]")
	{
		var (
			scannerBuf = make([]byte, 1024*1024)
			scanner    = bufio.NewScanner(buf)
		)
		scanner.Buffer(scannerBuf, cap(scannerBuf))
		for scanner.Scan() {
			fmt.Fprintln(io.Discard, scanner.Text())
		}

		fmt.Printf("\t結果： %v\n", scanner.Err())
	}

	return nil

	/*
		$ task
		task: [build] go build -o "/workspace/try-golang/try-golang" .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: scannerop_buffer

		[Name] "scannerop_buffer"
		一行のバイトサイズ: 65549bytes

		[デフォルトバッファサイズで読み込み]
		        結果： bufio.Scanner: token too long
		[拡張バッファサイズで読み込み]
		        結果： <nil>


		[Elapsed] 587.68µs
	*/
}
