package leak

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
	m["goroutines_leak_forgotten_sender"] = ForgottenSender
	m["goroutines_leak_forgotten_receiver"] = ForgottenReceiver
	m["goroutines_leak_abandoned_sender"] = AbandonedSender
	m["goroutines_leak_abandoned_receiver"] = AbandonedReceiver
	m["goroutines_leak_sender_after_error_check"] = SenderAfterErrorCheck
}
