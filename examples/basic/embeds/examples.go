package embeds

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
	m["embed_string"] = EmbedString
	m["embed_bytes"] = EmbedBytes
	m["embed_fs_singlefile"] = EmbedFsSingleFile
	m["embed_fs_multifiles"] = EmbedFsMultifiles
	m["embed_fs_filter"] = EmbedFsFilter
}
