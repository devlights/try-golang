package byteop

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["byteop_reader_from_byteslice"] = ReaderFromByteSlice
	m["byteop_cut_prefix_suffix"] = CutPrefixSuffix
	m["byteop_using_repeat"] = UsingRepeat
	m["byteop_go124_bytes_lines"] = Go124BytesLines
}
