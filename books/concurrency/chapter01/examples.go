package chapter01

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
	m["books_concurrency_ch01_racecondition"] = RaceCondition
	m["books_concurrency_ch01_racecondition_fix_with_mutex"] = RaceConditionFixWithMutex
	m["books_concurrency_ch01_racecondition_fix_with_channel"] = RaceConditionFixWithChannel
}
