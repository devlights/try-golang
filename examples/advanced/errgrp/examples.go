package errgrp

import (
	"github.com/devlights/try-golang/examples/advanced/errgrp/cmpwaitgroup"
	"github.com/devlights/try-golang/examples/advanced/errgrp/pipeline"
	"github.com/devlights/try-golang/examples/advanced/errgrp/withcontext"
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return &register{}
}

// Regist -- サンプルを登録します。
func (r *register) Regist(m mapping.ExampleMapping) {
	m["errgrp_error_with_waitgroup"] = cmpwaitgroup.ErrWithWaitGroup
	m["errgrp_error_with_errgroup"] = cmpwaitgroup.ErrWithErrGroup
	m["errgrp_with_context"] = withcontext.ErrGroupWithContext
	m["errgrp_with_pipeline"] = pipeline.ErrGroupWithPipeline
}
