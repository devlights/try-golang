package zeromemorycopy

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"github.com/devlights/gomy/times"
)

// StringToByteSlice -- 文字列からバイトスライスへメモリコピー無しに変換するサンプルです。
//
// REFERENCES
//   - https://stackoverflow.com/questions/59209493/how-to-use-unsafe-get-a-byte-slice-from-a-string-without-memory-copy
//   - https://github.com/devlights/gomy/blob/master/zeromemcpy/s2b.go
//   - https://pkg.go.dev/unsafe@go1.18.4#Slice
func StringToByteSlice() error {
	//
	// string から []byte へメモリコピー無しに変換するには unsafe パッケージを使う必要がある。
	// unsafeパッケージは文字通り unsafe な操作を行うパッケージなので、通常時には利用するべきではない。
	// パフォーマンスが極端に求められている場合で、且つ、メモリコピーの部分がボトルネックな場合にのみ利用するべき。
	// (通常、このような部分よりも他の部分がボトルネックになっているはず)
	//
	// Goは、基本的にこのようなトリッキーなことをしなくても充分速い。
	//

	// -------------------------------------
	// 大きなサイズの文字列を作る
	// -------------------------------------
	var sb strings.Builder
	for i := 0; i < 30_000_000; i++ {
		sb.WriteString(strconv.Itoa(i))
	}

	s := sb.String()
	fmt.Printf("[length] %vbytes, %vmb\n", len(s), len(s)/1024/1024)

	// -------------------------------------
	// 普通に変換
	// -------------------------------------
	elapsed := times.Stopwatch(func(start time.Time) {
		io.Discard.Write([]byte(s))
	})
	fmt.Printf("[normal] %v\n", elapsed)

	// -------------------------------------
	// メモリコピー無しで変換
	// -------------------------------------
	elapsed = times.Stopwatch(func(start time.Time) {
		io.Discard.Write(unsafe.Slice((*byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data)), len(s)))

		/* 上を細かく区切ると以下のようになる
		var (
			ptrStr     = unsafe.Pointer(&s)
			strHeader  = (*reflect.StringHeader)(ptrStr)
			ptrStrData = unsafe.Pointer(strHeader.Data)
			ptrByte    = (*byte)(ptrStrData)
			slice      = unsafe.Slice(ptrByte, len(s))
		)
		io.Discard.Write(slice)
		*/
	})
	fmt.Printf("[zeromemcpy] %v\n", elapsed)

	return nil
}
