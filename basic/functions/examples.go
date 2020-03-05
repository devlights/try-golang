package functions

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
	m["function_one_return_value"] = FunctionOneReturnValue
	m["function_multi_return_value"] = FunctionMultiReturnValue
	m["function_named_return_value"] = FunctionNamedReturnValue
}
