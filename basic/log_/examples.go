package log_

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
	m["log_flags"] = Flags
	m["log_prefix"] = Prefix
	m["log_sentry_basic"] = SentryBasic
	m["log_sentry_goroutine_bad"] = SentryGoroutineBad
	m["log_sentry_goroutine_good"] = SentryGoroutineGood
	m["log_output"] = Output
	m["log_new"] = NewLogger
}
