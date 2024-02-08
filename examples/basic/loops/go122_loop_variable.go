package loops

import "github.com/devlights/gomy/output"

// Go122LoopVariable は、Go 1.22 で導入された「ループ変数」の仕様変更についてのサンプルです.
//
// Go 1.22 より、for-range ループ中のループ変数がそれぞれ個別のアドレスを持つ変数となるように調整された。
//
// 以下、Go 1.22 のリリースノートより引用
//
// > Previously, the variables declared by a "for" loop were created once and updated by each iteration.
// In Go 1.22, each iteration of the loop creates new variables, to avoid accidental sharing bugs.
//
// > 以前は、"for "ループで宣言された変数は一度作成され、各反復で更新されていました。
// Go 1.22では、偶発的な共有バグを避けるため、ループの各反復で新しい変数が作成されます。
//
// # REFERENCES
//   - https://go.dev/blog/loopvar-preview
//   - https://go.dev/blog/go1.22
//   - https://go.dev/doc/go1.22#language
func Go122LoopVariable() error {
	type st struct {
		v int
	}

	var (
		items = []st{
			{v: 1},
			{v: 2},
			{v: 3},
		}
	)

	// Go 1.21 までは、同じアドレスが出力される
	// Go 1.22 からは、別々のアドレスが出力される
	for v := range items {
		output.Stdoutf("[Value (Addr)]", "%v (%p)\n", v, &v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: loops_go122_loop_variable

	   [Name] "loops_go122_loop_variable"
	   [Value (Addr)]       0 (0xc0001a2908)
	   [Value (Addr)]       1 (0xc0001a2910)
	   [Value (Addr)]       2 (0xc0001a2918)


	   [Elapsed] 37.92µs
	*/

}
