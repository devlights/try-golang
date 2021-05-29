package runner

import "github.com/devlights/try-golang/mapping"

type (
	SilentExample struct {
		Called bool
	}
	SilentRegister struct {
		Target *SilentExample
	}
)

func (me *SilentExample) Run() error {
	me.Called = true
	return nil
}

func (me *SilentRegister) Regist(m mapping.ExampleMapping) {
	e := &SilentExample{false}
	m["silent"] = e.Run

	me.Target = e
}
