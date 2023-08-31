package builtins

import "github.com/devlights/gomy/output"

// Delete は、ビルトイン関数 delete についてのサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/builtin@go1.21.0#delete
func Delete() error {
	//
	// delete は、mapから特定のキーを削除するビルトイン関数
	//   - 存在しないキーを指定してもエラーにはならない
	//
	var (
		m = map[string]int{
			"apple": 100,
			"ringo": 200,
		}
		p = func(m map[string]int) {
			output.Stdoutl("[map]", m)
		}
	)

	p(m)
	delete(m, "ringo")
	delete(m, "hoge")
	p(m)

	return nil
}
