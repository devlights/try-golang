package variables

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
	m["var_statement_declare"] = VarStatementDeclares
	m["package_scope_variable"] = PackageScopeVariable
	m["short_assignment_statement"] = ShortAssignmentStatement
	m["shadowing_variable"] = ShadowingVariable
}
