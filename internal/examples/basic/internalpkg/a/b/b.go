package b

import (
	"strings"

	"github.com/devlights/try-golang/examples/basic/internalpkg/internal/sub1"
)

// B -- サンプルから呼び出される関数
func B() string {
	return strings.ToUpper(sub1.CallSub1())
}
