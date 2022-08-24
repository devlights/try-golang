package formatting

import (
	"fmt"

	"github.com/devlights/gomy/output"
)

// AdverbAsterisk -- フォーマッティングの %*s についてのサンプルです
//
// フォーマット書式に出てくる * は、Explicit argument indexes の仕様の一つ。
// 詳細は、同ディレクトリにある formatting_adverb_explicit_argument_indexes のサンプル参照。
//
// %*s というのは、パラメータのポジション指定を省略している状態。
// なので、「%*s」というのは「%[1]*[2]s」と同じことになる。
//
// REFERENCES:
//   - 書籍 「プログラミング言語Go」 P.152
//   - https://golang.org/fmt/
func AdverbAsterisk() error {
	levels := []int{1, 2, 4, 8}

	for _, v := range levels {
		// %*s というのは、パラメータのポジション指定を省略している状態。
		// なので、「%*s」というのは「%[1]*[2]s」と同じことになる。
		//
		// 以下と同じことになる
		//   - %1s%s
		//   - %2s%s
		//   - %4s%s
		//   - %8s%s
		fmt.Printf("%*s%s\n", v, "", "message")
	}

	// 上記と同じ内容をポジション指定を付与すると以下となる
	output.StdoutHr()
	for _, v := range levels {
		// 一つ目のパラメータの値は * が付与されているので、そのまま値として書式に利用される
		// ２つ目と３つ目のパラメータの値は * が付与されていないので通常通り利用される
		fmt.Printf("%[1]*[2]s%[3]s\n", v, "", "message")
	}

	// なので、ゼロ埋めしたい場合は以下のようにすれば出来る
	output.StdoutHr()
	for _, v := range levels {
		fmt.Printf("%0*s%s\n", v, "", "message")
		// 以下と同じ
		// fmt.Printf("%0[1]*[2]s%[3]s\n", v, "", "message")
	}

	return nil
}
