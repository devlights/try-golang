package times

import (
	"github.com/devlights/try-golang/mappings"
)

type (
	register struct{}
)

func NewRegister() mappings.Register {
	return new(register)
}

func (r *register) Regist(m mappings.ExampleMapping) {
	m["time_since"] = TimeSince
	m["time_after"] = TimeAfter
	m["time_unix_to_time"] = TimeUnixToTime
	m["time_now"] = TimeNow
	m["time_parse"] = TimeParse
}
