package structs

import (
	"unsafe"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/internal/examples/basic/structs/types"
)

// MemoryPadding は、構造体メンバーの定義順によってGoランタイムがメモリ上にパディングを挿入することを確認するサンプルです.
func MemoryPadding() error {
	var (
		notGood = types.MemoryPadding{}
		good = types.NoMemoryPadding{}
	)

	output.Stdoutl("[Padding 発生]", unsafe.Sizeof(notGood))
	output.Stdoutl("", notGood.Layout())
	output.Stdoutl("[Padding なし]", unsafe.Sizeof(good))
	output.Stdoutl("", good.Layout())

	return nil
}