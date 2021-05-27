package builder

import (
	"github.com/devlights/try-golang/examples/advanced"
	"github.com/devlights/try-golang/examples/basic"
	"github.com/devlights/try-golang/examples/effectivego"
	"github.com/devlights/try-golang/examples/gotour"
	"github.com/devlights/try-golang/mapping"
)

// BuildMappings は、サンプル実行のためのマッピング情報を構築します.
func BuildMappings() mapping.ExampleMapping {
	m := mapping.NewSampleMapping()

	m.MakeMapping(
		advanced.NewRegister(),
		basic.NewRegister(),
		effectivego.NewRegister(),
		gotour.NewRegister(),
	)

	return m
}
