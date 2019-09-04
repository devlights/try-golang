// Go の プログラムは、パッケージで構成される。
// 規約により、パッケージ名はimportパスの最後の要素の名前となる。
// プログラムは、必ず main パッケージから開始される。
// Go では、一つのディレクトリ内に一つのパッケージしか含めることが出来ない。
package tutorial

import (
	"fmt"
)

// GoTourHelloWorld は、 [A Tour of Go](http://bit.ly/2HsCMiG) の 要約.
func GoTourHelloWorld() error {
	// ------------------------------------------------------------
	// Hello World
	//   文字列を出力するには、 fmt パッケージの Println() などを使う
	// ------------------------------------------------------------
	fmt.Println("Hello World")

	return nil
}
