package errs

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	return new(register)
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["error_basic"] = Basic
	m["error_sentinel"] = Sentinel
	m["error_typeassertion"] = TypeAssertion
	m["error_wrap_unwrap"] = WrapAndUnwrap
}
