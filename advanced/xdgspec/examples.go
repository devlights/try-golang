package xdgspec

import "github.com/devlights/try-golang/interfaces"

type (
	register struct{}
)

func NewRegister() interfaces.Register {
	r := new(register)
	return r
}

func (r *register) Regist(m interfaces.ExampleMapping) {
	m["xdg_base_directory"] = XdgBaseDirectory
	m["xdg_user_directory"] = XdgUserDirectory
	m["xdg_file_operation"] = XdgFileOperation
}
