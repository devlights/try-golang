package internalpkg

import (
	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/basic/internalpkg/a/b"
	"github.com/devlights/try-golang/examples/basic/internalpkg/internal/sub1"
	// Compile Error: use of internal package github.com/devlights/try-golang/examples/basic/internalpkg/internal/internal/sub2 not allowed
	// "github.com/devlights/try-golang/examples/basic/internalpkg/internal/internal/sub2"
)

// Basic -- Go 1.14 から導入された internal packages の挙動を確認するサンプルです.
//
// REFERENCES:
//   - https://qiita.com/rema424/items/2dc22ef36ab6aba64e20
//   - https://golang.org/doc/go1.4#internalpackages
func Basic() error {
	// -------------------------------------------------------
	// Go 1.14 から internal という名前のパッケージ名は
	// 特別扱いされるようになった。
	//
	// internal というパッケージが存在する場合
	// internal パッケージと同じ階層から辿ることが可能な
	// 場所ではインポートが可能になる。
	//
	// つまり、intenalがいる場所を起点として辿ることが
	// 可能であれば、インポート可能。その途中に internal が
	// 出てきたら、次は、そのinternalを起点として辿ることが
	// 可能な場合はインポート可能。上に上がらないと辿れない
	// 場合はインポート不可となる。
	//
	// 本サンプルでは、以下の2階層のinternalを配置している.
	//
	// internalpkg
	//   - internalpkg_basic.go      (1)
	//   - internal
	//     - sub1
	//       - internalpkg_sub1.go   (2)
	//     - internal
	//       - sub2
	//         - internalpkg_sub2.go (3)
	//   - a
	//     - b
	//       - b.go                  (4)
	//
	// (1) から (2) はインポート可能だが、(3)はインポート不可
	// (2) から (3) はインポート可能
	// (4) から (2) はインポート可能だが、(3)はインポート不可
	// -------------------------------------------------------

	// (1) から (2) はインポート可能
	output.Stdoutl("[From (1) to (2)]", sub1.CallSub1())

	// (1) から (3) はインポート不可
	// 兄弟階層ではない internal パッケージはインポートできない
	//
	// Compile Error: use of internal package github.com/devlights/try-golang/examples/basic/internalpkg/internal/internal/sub2 not allowed
	// output.Stdoutl("[sub2]", sub2.InternalPkgSub2())

	// (2) から (3) はインポート可能
	output.Stdoutl("[From (2) to (3)]", sub1.CallSub2())

	// (4) から (2) はインポート可能だが、(3)はインポート不可
	output.Stdoutl("[From (4) to (2)]", b.B())

	return nil
}
