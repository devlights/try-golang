package errs

import (
	"github.com/devlights/try-golang/mappings"
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
	m["error_basic"] = Basic
	m["error_sentinel"] = Sentinel
	m["error_typeassertion"] = TypeAssertion
	m["error_wrap_unwrap"] = WrapAndUnwrap
}
