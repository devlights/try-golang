package lib

import (
	"github.com/devlights/try-golang/advanced"
	"github.com/devlights/try-golang/basic"
	"github.com/devlights/try-golang/books"
	"github.com/devlights/try-golang/effectivego"
	"github.com/devlights/try-golang/interfaces"
	"github.com/devlights/try-golang/tutorial"
)

// MakeMapping は、サンプル実行のためのマッピング情報を生成します.
func MakeMapping() interfaces.ExampleMapping {
	mapping := interfaces.NewSampleMapping()

	mapping.MakeMapping(
		advanced.NewRegister(),
		basic.NewRegister(),
		books.NewRegister(),
		effectivego.NewRegister(),
		tutorial.NewRegister(),
	)

	return mapping
}
