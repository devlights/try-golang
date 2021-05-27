package loops

import "github.com/devlights/gomy/output"

// BasicForLoop は、他の言語でも同じように存在する基本的な for-loop についてのサンプルです.
func BasicForLoop() error {
	var (
		items = []string{
			"go",
			"java",
			"dotnet",
			"python",
			"flutter",
		}
	)

	// 他の言語と同じように Go にも インデックス 付きの for-loop がある
	for i := 0; i < 5; i++ {
		output.Stdoutf("", "[%d] %s\n", i, items[i])
	}

	return nil
}
