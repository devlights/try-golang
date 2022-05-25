package exp_constraints

import (
	"github.com/devlights/gomy/output"
	"golang.org/x/exp/constraints"
)

type myStr string

func add[E constraints.Signed](x, y E) E {
	return x + y
}

func less[E constraints.Ordered](x, y E) bool {
	return x < y
}

// Constraints -- Go 1.18 リリース時には含まれなかった制約型が定義されている golang.org/x/exp/constraints パッケージのサンプルです。
func Constraints() error {
	//
	// constraints パッケージには、Go 1.18 リリース時には
	// 含まれなかった制約型が定義されている。
	//
	// - Complex
	// - Float
	// - Integer
	// - Ordered
	// - Signed
	// - Unsigned
	//
	// 特に Ordered は助かる。
	//
	output.Stdoutl("[Signed]", add(int32(1), int32(3)))
	output.Stdoutl("[Signed]", add(int8(1), int8(-2)))

	output.Stdoutl("[Ordered]", less(1, 2))
	output.Stdoutl("[Ordered]", less("z", "b"))
	output.Stdoutl("[Ordered]", less(myStr("b"), myStr("z")))

	return nil
}
