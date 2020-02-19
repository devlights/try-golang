package interfaces

type (
	// SampleKeyは、サンプル名を表すキーを表します
	ExampleKey string

	// SampleFuncは、実行するサンプル処理を表します
	ExampleFunc func() error

	// SampleMappingは、サンプルのマッピング定義を表します
	ExampleMapping map[ExampleKey]ExampleFunc
)

// NewSampleMapping は、SampleMappingのコンストラクタ関数です
func NewSampleMapping() ExampleMapping {
	return make(ExampleMapping)
}

// MakeMapping は、マッピングを生成します
func (m ExampleMapping) MakeMapping(registers ...Register) {
	for _, register := range registers {
		register.Regist(m)
	}
}
