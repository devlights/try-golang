package filepaths

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
	m["filepath_walk"] = FilePathWalk
}
