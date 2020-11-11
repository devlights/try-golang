package builder

import (
	"github.com/devlights/try-golang/internal/effectivego"
	"github.com/devlights/try-golang/internal/examples/advanced"
	"github.com/devlights/try-golang/internal/examples/basic"
	"github.com/devlights/try-golang/pkg/mappings"
	"github.com/devlights/try-golang/internal/tutorial"
)

// BuildMappings は、サンプル実行のためのマッピング情報を構築します.
func BuildMappings() mappings.ExampleMapping {
	m := mappings.NewSampleMapping()

	m.MakeMapping(
		advanced.NewRegister(),
		basic.NewRegister(),
		effectivego.NewRegister(),
		tutorial.NewRegister(),
	)

	return m
}
