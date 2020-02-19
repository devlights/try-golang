package interfaces

type (
	// SampleKeyは、サンプル名を表すキーを表します
	SampleKey string

	// SampleFuncは、実行するサンプル処理を表します
	SampleFunc func() error

	// SampleMappingは、サンプルのマッピング定義を表します
	SampleMapping map[SampleKey]SampleFunc
)

// NewSampleMapping は、SampleMappingのコンストラクタ関数です
func NewSampleMapping() SampleMapping {
	return make(SampleMapping)
}

// MakeMapping は、マッピングを生成します
func (m SampleMapping) MakeMapping(registers ...Register) {
	for _, register := range registers {
		register.Regist(m)
	}
}
