package xdgspec

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["xdg_base_directory"] = XdgBaseDirectory
	m["xdg_user_directory"] = XdgUserDirectory
	m["xdg_file_operation"] = XdgFileOperation
}
