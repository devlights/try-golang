package embeds

import (
	"github.com/devlights/try-golang/pkg/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["embed_string"] = EmbedString
	m["embed_bytes"] = EmbedBytes
	m["embed_fs_singlefile"] = EmbedFsSingleFile
	m["embed_fs_multifiles"] = EmbedFsMultifiles
	m["embed_fs_filter"] = EmbedFsFilter
}
