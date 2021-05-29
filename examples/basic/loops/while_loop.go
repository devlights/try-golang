package loops

import (
	"github.com/devlights/gomy/output"
)

// WhileLoop は、GoでのWhileループについてのサンプルです.
func WhileLoop() error {
	// Go には、ループはすべて for で記載することになっている。
	// 他の言語にある while () {} は提供されていない。
	count := 5
	for count > 0 {
		output.Stdoutl("[count]", count)
		count -= 1
	}

	return nil
}
