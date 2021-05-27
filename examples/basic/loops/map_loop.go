package loops

import "github.com/devlights/gomy/output"

// MapLoop は、 map のループについてのサンプルです.
func MapLoop() error {
	var (
		m = map[string]string{
			"go":     "fmt.Println",
			"java":   "System.out.println",
			"dotnet": "Console.WriteLine",
			"python": "print",
		}
	)

	// map の　ループ は、key, value の値が毎ターン取得できる
	for k, v := range m {
		output.Stdoutf("", "[%s] %s\n", k, v)
	}

	return nil
}
