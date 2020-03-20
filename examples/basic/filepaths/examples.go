package filepaths

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return &register{}
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["filepath_walk"] = FilePathWalk
	m["filepath_glob"] = FilePathGlob
}
