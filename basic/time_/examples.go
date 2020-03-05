package time_

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
	m["time_since"] = TimeSince
	m["time_after"] = TimeAfter
	m["time_unix_to_time"] = TimeUnixToTime
	m["time_now"] = TimeNow
	m["time_parse"] = TimeParse
}
