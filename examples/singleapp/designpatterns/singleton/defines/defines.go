package defines

import "sync"

type (
	Defines struct {
		def1 string
		def2 string
	}
)

func (me Defines) GetDef1() string { return me.def1 }
func (me Defines) GetDef2() string { return me.def2 }

var (
	instance Defines
	once     sync.Once
)

func GetInstance() Defines {
	once.Do(func() {
		instance = Defines{
			def1: "hello",
			def2: "world",
		}

		println("create: Defines")
	})

	return instance
}
