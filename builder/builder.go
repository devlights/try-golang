package builder

import (
	"github.com/devlights/try-golang/books"
	"github.com/devlights/try-golang/effectivego"
	"github.com/devlights/try-golang/examples/advanced"
	"github.com/devlights/try-golang/examples/basic"
	"github.com/devlights/try-golang/mappings"
	"github.com/devlights/try-golang/tutorial"
)

// BuildMappings は、サンプル実行のためのマッピング情報を構築します.
func BuildMappings() mappings.ExampleMapping {
	m := mappings.NewSampleMapping()

	m.MakeMapping(
		advanced.NewRegister(),
		basic.NewRegister(),
		books.NewRegister(),
		effectivego.NewRegister(),
		tutorial.NewRegister(),
	)

	return m
}
