package structs

import (
	"github.com/devlights/gomy/output"
)

// BlankIdentifier -- 構造体定義時に blank identifier を意図的に用意して初期化時にフィールド名の指定を必須にするやり方のサンプルです.
//
// REFERENCES:
//   - https://qiita.com/fuubit/items/88ff1185de1a67d9e5bd
func BlankIdentifier() error {
	type (
		forceNamedParamInit struct {
			Value string
			_     struct{} // blank identifier を配置すると、外部から初期化する際に常にフィールド名付きで初期化を必須にできる
		}

		nonForceNamedParamInit struct {
			Value string
		}
	)

	// 空のフィールドが存在する構造体の場合は、フィールド名の指定をしないと初期化できない
	o1 := forceNamedParamInit{Value: "ok pattern"}
	// 以下はコンパイルできない (Too few values)
	// o2 := forceNamedParamInit{"ng pattern"}

	// 空のフィールドがない場合は、フィールド名を指定してもしなくても初期化できる
	o3 := nonForceNamedParamInit{Value: "ok pattern"}
	o4 := nonForceNamedParamInit{"ok pattern"}

	output.Stdoutl("result", o1, o3, o4)

	return nil
}
