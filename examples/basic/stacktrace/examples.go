package stacktrace

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["stacktrace_debug_printstack"] = DebugPrintStack
	m["stacktrace_debug_stack"] = DebugStack
	m["stacktrace_runtime_stack"] = RuntimeStack
	m["stacktrace_pprof_writeto"] = PprofGoroutineWriteTo
}
