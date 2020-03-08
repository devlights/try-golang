package logging

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
	m["log_flags"] = Flags
	m["log_prefix"] = Prefix
	m["log_sentry_basic"] = SentryBasic
	m["log_sentry_goroutine_bad"] = SentryGoroutineBad
	m["log_sentry_goroutine_good"] = SentryGoroutineGood
	m["log_output"] = Output
	m["log_new"] = NewLogger
}
