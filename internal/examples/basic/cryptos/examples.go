package cryptos

import (
	"github.com/devlights/try-golang/internal/examples/basic/cryptos/checksum"
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
	m["crypto_md5_checksum"] = checksum.Md5Checksum
}
