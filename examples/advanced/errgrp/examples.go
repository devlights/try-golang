package errgrp

import (
	"github.com/devlights/try-golang/examples/advanced/errgrp/cmpwaitgroup"
	"github.com/devlights/try-golang/examples/advanced/errgrp/pipeline"
	"github.com/devlights/try-golang/examples/advanced/errgrp/withcontext"
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return &register{}
}

// Regist -- サンプルを登録します。
func (r *register) Regist(m mappings.ExampleMapping) {
	m["errgrp_error_with_waitgroup"] = cmpwaitgroup.ErrWithWaitGroup
	m["errgrp_error_with_errgroup"] = cmpwaitgroup.ErrWithErrGroup
	m["errgrp_with_context"] = withcontext.ErrGroupWithContext
	m["errgrp_with_pipeline"] = pipeline.ErrGroupWithPipeline
}
