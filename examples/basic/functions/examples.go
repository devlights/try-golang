package functions

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
	m["function_one_return_value"] = FunctionOneReturnValue
	m["function_multi_return_value"] = FunctionMultiReturnValue
	m["function_named_return_value"] = FunctionNamedReturnValue
}