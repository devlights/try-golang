package japanese

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	return &register{}
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["sjis_readwrite"] = SjisReadWrite
	m["eucjp_readwrite"] = EucJpReadWrite
	m["gomy_readwrite"] = GomyReadWrite
}
