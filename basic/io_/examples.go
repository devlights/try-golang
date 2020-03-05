package io_

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
	m["fileio01"] = FileIo01
	m["fileio02"] = FileIo02
	m["fileio03"] = FileIo03
	m["fileio04"] = FileIo04

}
