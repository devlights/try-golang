package loops

import "github.com/devlights/gomy/output"

// BasicForeach は、Go での foreach ループについてのサンプルです.
func BasicForeach() error {
	var (
		items = []string{
			"go",
			"java",
			"dotnet",
			"python",
			"flutter",
		}
	)

	for i, v := range items {
		output.Stdoutf("", "[%d] %s\n", i, v)
	}

	return nil
}
