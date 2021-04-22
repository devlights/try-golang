package structs

import (
	"unsafe"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/internal/examples/basic/structs/types"
)

// MemoryPadding は、構造体メンバーの定義順によってGoランタイムがメモリ上にパディングを挿入することを確認するサンプルです.
func MemoryPadding() error {
	var (
		st4bytes = types.Struct4Bytes{}    // メモリ上のサイズが 4bytes になる構造体
		st8bytes = types.Struct8Bytes{}    // メモリ上のサイズが 8bytes になる構造体
		notGood  = types.MemoryPadding{}   // メモリのパディングが発生する構造体
		good     = types.NoMemoryPadding{} // メモリのパディングが発生しない構造体
	)

	output.Stdoutf("[st4bytes]", "%d byte(s)\n", unsafe.Sizeof(st4bytes))
	output.StdoutHr()

	output.Stdoutf("[st8bytes]", "%d byte(s)\n", unsafe.Sizeof(st8bytes))
	output.StdoutHr()

	output.Stdoutf("[Padding 発生]", "%d byte(s)\n", unsafe.Sizeof(notGood))
	output.Stdoutl("", notGood.Layout())
	output.StdoutHr()

	output.Stdoutf("[Padding なし]", "%d byte(s)\n", unsafe.Sizeof(good))
	output.Stdoutl("", good.Layout())

	return nil
}
