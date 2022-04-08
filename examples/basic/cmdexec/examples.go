package cmdexec

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
	m["cmdexec_oneshot"] = OneShot
	m["cmdexec_oneshot_with_stderr"] = OneShotWithStderr
	m["cmdexec_stdinouterr"] = Stdinouterr
	m["cmdexec_withcontext"] = WithContext
	m["cmdexec_pipe"] = Pipe
	m["cmdexec_multi_command_with_pipe"] = MultiCommandWithPipe
	m["cmdexec_env"] = WithEnv
}
