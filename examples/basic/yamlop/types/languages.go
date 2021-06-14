package types

type (
	Version struct {
		Major int `yaml:"major"`
		Minor int `yaml:"minor"`
	}

	Language struct {
		Name    string  `yaml:"name"`
		PrintFn string  `yaml:"printfn"`
		Version Version `yaml:"version"`
	}

	YamlData struct {
		Languages []Language `yaml:"languages"`
	}
)
