package times

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
	m["time_since"] = TimeSince
	m["time_sub"] = TimeSub
	m["time_after"] = TimeAfter
	m["time_unix_to_time"] = TimeUnixToTime
	m["time_now"] = TimeNow
	m["time_parse"] = TimeParse
	m["time_tick_and_ticker"] = TickAndTicker
	m["time_timer"] = Timer
	m["time_afterfunc"] = AfterFunc
	m["time_changelocaltimezone"] = ChangeLocalTimezone
	m["time_truncate"] = Truncate
	m["time_json"] = TimeJson
}
