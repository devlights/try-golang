package variables

import (
	"github.com/devlights/try-golang/examples/basic/variables/loopiterator"
	"github.com/devlights/try-golang/examples/basic/variables/packagescope"
	"github.com/devlights/try-golang/examples/basic/variables/shadowing"
	"github.com/devlights/try-golang/examples/basic/variables/shortassignment"
	"github.com/devlights/try-golang/examples/basic/variables/varstatement"
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	return new(register)
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["var_statement_declare"] = varstatement.Basic
	m["package_scope_variable"] = packagescope.Basic
	m["short_assignment_statement"] = shortassignment.Basic
	m["shadowing_variable"] = shadowing.Basic
	m["using_ref_to_loop_iterator_variable"] = loopiterator.CommonMistakePattern
}
