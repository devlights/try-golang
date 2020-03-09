package errgrp

import (
	"github.com/devlights/try-golang/examples/advanced/errgrp/cmpwaitgroup"
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	return &register{}
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["errgrp_error_with_waitgroup"] = cmpwaitgroup.ErrWithWaitGroup
	m["errgrp_error_with_errgroup"] = cmpwaitgroup.ErrWithErrGroup
}
