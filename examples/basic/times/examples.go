package times

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mappings.Register を生成します。
func NewRegister() mappings.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mappings.ExampleMapping) {
	m["time_since"] = TimeSince
	m["time_after"] = TimeAfter
	m["time_unix_to_time"] = TimeUnixToTime
	m["time_now"] = TimeNow
	m["time_parse"] = TimeParse
}
