package structs

import (
	"unsafe"

	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/structs/types"
)

// MemoryPadding は、構造体メンバーの定義順によってGoランタイムがメモリ上にパディングを挿入することを確認するサンプルです.
//
// REFERENCES:
//   - https://itnext.io/structure-size-optimization-in-golang-alignment-padding-more-effective-memory-layout-linters-fffdcba27c61
//   - https://logicalbeat.jp/blog/4032/
//   - https://ja.wikipedia.org/wiki/%E3%83%87%E3%83%BC%E3%82%BF%E6%A7%8B%E9%80%A0%E3%82%A2%E3%83%A9%E3%82%A4%E3%83%A1%E3%83%B3%E3%83%88
func MemoryPadding() error {
	var (
		st4bytes = types.Struct4Bytes{}    // メモリ上のサイズが 4bytes になる構造体
		st8bytes = types.Struct8Bytes{}    // メモリ上のサイズが 8bytes になる構造体
		notGood  = types.MemoryPadding{}   // メモリのパディングが発生する構造体
		good     = types.NoMemoryPadding{} // メモリのパディングが発生しない構造体
	)

	output.Stdoutf("[st4bytes]", "%d byte(s)\n", unsafe.Sizeof(st4bytes))
	output.Stdoutl("", st4bytes.Layout())
	output.StdoutHr()

	output.Stdoutf("[st8bytes]", "%d byte(s)\n", unsafe.Sizeof(st8bytes))
	output.Stdoutl("", st8bytes.Layout())
	output.StdoutHr()

	output.Stdoutf("[Padding 発生]", "%d byte(s)\n", unsafe.Sizeof(notGood))
	output.Stdoutl("", notGood.Layout())
	output.StdoutHr()

	output.Stdoutf("[Padding なし]", "%d byte(s)\n", unsafe.Sizeof(good))
	output.Stdoutl("", good.Layout())

	return nil

	/*
		$ task
		task: [build] go build .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: struct_memory_padding

		[Name] "struct_memory_padding"
		[st4bytes]           4 byte(s)

		        | Flag     |             | Value     |
		        --------------------------------------
		        | bool (1) | padding (2) | int16 (2) |
		        |                  4                 |

		--------------------------------------------------
		[st8bytes]           8 byte(s)

		        | Flag     |             | Value     |
		        --------------------------------------
		        | bool (1) | padding (3) | int32 (4) |
		        |          4             |     4     |

		--------------------------------------------------
		[Padding 発生]         12 byte(s)

		        | Flag1    |             | ShortVal  | Flag2    |             | FloatVal    |
		        -----------------------------------------------------------------------------
		        | bool (1) | padding (1) | int16 (2) | bool (1) | padding (3) | float32 (4) |
		        |                  4                 |          4             |      4      |

		--------------------------------------------------
		[Padding なし]         8 byte(s)

		        | FloatVal    | ShortVal  | Flag1    | Flag2    |
		        -------------------------------------------------
		        | float32 (4) | int16 (2) | bool (1) | bool (1) |
		        |     4       |              4                  |



		[Elapsed] 96.28µs
	*/

}
