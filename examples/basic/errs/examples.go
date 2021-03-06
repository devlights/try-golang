package errs

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
	m["error_basic"] = Basic
	m["error_sentinel"] = Sentinel
	m["error_typeassertion"] = TypeAssertion
	m["error_wrap_unwrap"] = WrapAndUnwrap
}
