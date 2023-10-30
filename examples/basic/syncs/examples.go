package syncs

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
	m["syncs_no_sync"] = NoSync
	m["syncs_atomic_types"] = AtomicTypes
	m["syncs_atomic_add"] = AtomicAdd
	m["syncs_atomic_compare_and_swap"] = CompareAndSwap
	m["syncs_use_channel"] = UseChannel
	m["syncs_use_mutex"] = UseMutex
	m["syncs_use_cond_signal"] = UseCondSignal
	m["syncs_use_cond_broadcast"] = UseCondBroadcast
	m["syncs_map"] = Map
}
