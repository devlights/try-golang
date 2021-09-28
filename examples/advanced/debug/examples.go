package debug

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
//
// REFERENCES:
//   - https://yellow-peacock-3ct2cjzl.ws-us18.gitpod.io/
func NewRegister() mapping.Register {
	return &register{}
}

// Regist -- サンプルを登録します。
func (r *register) Regist(m mapping.ExampleMapping) {
	m["debug_build_info"] = BuildInfo
}
