package filepaths

import (
	"github.com/devlights/try-golang/interfaces"
)

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	return &register{}
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["filepath_walk"] = FilePathWalk
}
