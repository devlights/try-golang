package effectivego23

import (
	"github.com/devlights/gomy/output"
)

type (
	myInterface interface {
		run()
	}

	myImpl struct {
	}
)

func (m *myImpl) run() {
	output.Stdoutl("helloworld")
}

// *myImpl は myInterface を実装する
// この変換チェックはコンパイル時に走るので、インターフェースが変更された場合
// このパッケージはコンパイル出来ない状態となる
var _ myInterface = (*myImpl)(nil)

// InterfaceCheck -- Effective Go - Interface checks の 内容についてのサンプルです。
func InterfaceCheck() error {
	/*
		https://golang.org/doc/effective_go.html#blank_implements

		Goではインターフェースを実装するのに明示的な宣言が必要なく
		単純に特定のインターフェースが備えるメソッドを全て実装していれば良い。

		ただし、弊害としてインターフェースに変更が加わった場合に、いつの間にか
		その型がインターフェースの実装できていない状態になってしまうことがたまにある。

		通常は、IDEなどを利用していると教えてくれるものであるが
		明示的に型がインターフェースの実装を出来ているのかをチェックする方法がある。

		パッケージレベルの変数として以下の変数宣言を行うことにより、その型がインターフェースを
		実装出来ていない場合はコンパイルエラーになる。

			var _ インターフェース = (型)(nil)

		便利ではあるが、このチェックは全ての型に対して行うべきではないとEffective Go には記載されている。

		> ただし、インターフェイスを満たすすべてのタイプに対してこれを実行しないでください。
		> 慣例により、このような宣言は、コードに既に静的な変換が存在しない場合にのみ使用されます。
		> これはまれなイベントです。
	*/
	return nil
}
