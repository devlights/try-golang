package loops

import "github.com/devlights/gomy/output"

// SliceLoop は、スライスのループについてのサンプルです.
func SliceLoop() error {
	var (
		items = []string{
			"golang",
			"java",
			"dotnet",
			"python",
		}
	)

	// スライスの foreach は、インデックスと値 となる
	for i, v := range items {
		output.Stdoutf("", "[%d] %s\n", i, v)
	}

	return nil
}
