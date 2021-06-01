package fileio

import (
	"github.com/devlights/try-golang/examples/basic/fileio/filesystem"
	"github.com/devlights/try-golang/examples/basic/fileio/readwrite"
	"github.com/devlights/try-golang/examples/basic/fileio/stat"
	"github.com/devlights/try-golang/examples/basic/fileio/stdinouterr"
	"github.com/devlights/try-golang/mapping"
)

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	readwrite.NewRegister().Regist(m)
	stat.NewRegister().Regist(m)
	stdinouterr.NewRegister().Regist(m)
	filesystem.NewRegister().Regist(m)
}
