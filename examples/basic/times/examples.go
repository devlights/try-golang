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
	m["time_truncate_hours"] = TruncateHours
	m["time_json"] = TimeJson
	m["time_json_custom"] = TimeJsonCustom
	m["time_in"] = TimeIn
	m["time_change_timezone"] = ChangeTimeZone
	m["time_format_datetime"] = FormatDateTime
	m["time_format_dateonly"] = FormatDateOnly
	m["time_format_timeonly"] = FormatTimeOnly
	m["time_format_millisecond"] = FormatMillisecond
	m["time_format_microsecond"] = FormatMicrosecond
	m["time_calc_nextmonth"] = CalcNextMonth
	m["time_daysinmonth"] = DaysInMonth
	m["time_sleep"] = Sleep
	m["time_cancellable_sleep"] = CancellableSleep
	m["time_parse_duration"] = ParseDuration
	m["time_do_n_duration"] = DoNDurations
}
