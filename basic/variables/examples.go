package variables

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
	m["var_statement_declare"] = VarStatementDeclares
	m["package_scope_variable"] = PackageScopeVariable
	m["short_assignment_statement"] = ShortAssignmentStatement
	m["shadowing_variable"] = ShadowingVariable
}
