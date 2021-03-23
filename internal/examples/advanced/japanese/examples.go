package japanese

import (
	"github.com/devlights/try-golang/pkg/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return &register{}
}

// Regist -- サンプルを登録します。
func (r *register) Regist(m mappings.ExampleMapping) {
	m["sjis_readwrite"] = SjisReadWrite
	m["eucjp_readwrite"] = EucJpReadWrite
}
