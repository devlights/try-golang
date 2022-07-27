package zeromemorycopy

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"
	"unsafe"

	"github.com/devlights/gomy/times"
	"github.com/devlights/gomy/zeromemcpy"
)

// ByteSliceToString -- バイトスライスから文字列へメモリコピー無しに変換するサンプルです。
//
// REFERENCES
//   - https://github.com/devlights/gomy/blob/master/zeromemcpy/b2s.go
//   - https://cs.opensource.google/go/go/+/refs/tags/go1.18.4:src/strings/builder.go;l=47
func ByteSliceToString() error {
	//
	// []byte から string へメモリコピー無しに変換するには unsafe パッケージを使う必要がある。
	// unsafeパッケージは文字通り unsafe な操作を行うパッケージなので、通常時には利用するべきではない。
	// パフォーマンスが極端に求められている場合で、且つ、メモリコピーの部分がボトルネックな場合にのみ利用するべき。
	// (通常、このような部分よりも他の部分がボトルネックになっているはず)
	//
	// Goは、基本的にこのようなトリッキーなことをしなくても充分速い。
	//

	// -------------------------------------
	// 大きなサイズのバッファを作る
	// -------------------------------------

	buf := new(bytes.Buffer)
	for i := 0; i < 30_000_000; i++ {
		buf.WriteString(strconv.Itoa(i))
	}

	b := buf.Bytes()
	fmt.Printf("[length] %vbytes, %vmb\n", len(b), len(b)/1024/1024)

	// -------------------------------------
	// 普通に変換
	// -------------------------------------
	elapsed := times.Stopwatch(func(start time.Time) {
		io.WriteString(io.Discard, string(b))
	})
	fmt.Printf("[normal] %v\n", elapsed)

	// -------------------------------------
	// メモリコピー無しで変換
	// -------------------------------------
	elapsed = times.Stopwatch(func(start time.Time) {
		io.WriteString(io.Discard, *(*string)(unsafe.Pointer(&b)))

		/* 上を細かく区切ると以下のようになる
		var (
			ptrByte = unsafe.Pointer(&b)
			ptrStr  = (*string)(ptrByte)
			str     = *ptrStr
		)
		io.WriteString(io.Discard, str)
		*/
	})
	fmt.Printf("[zeromemcpy] %v\n", elapsed)

	// -------------------------------------
	// zeromemcpy.b2s
	// -------------------------------------
	elapsed = times.Stopwatch(func(start time.Time) {
		io.WriteString(io.Discard, zeromemcpy.B2S(b))
	})
	fmt.Printf("[zeromemcpy.B2S] %v\n", elapsed)

	return nil
}
