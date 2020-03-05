package error_

import (
	"github.com/devlights/try-golang/interfaces"
)

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	return new(register)
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["error_basic"] = Basic
	m["error_sentinel"] = Sentinel
	m["error_typeassertion"] = TypeAssertion
	m["error_wrap_unwrap"] = WrapAndUnwrap
}
