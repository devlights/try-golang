package maths

import (
	"fmt"
	"math/rand/v2"
	"os"
	"text/tabwriter"
)

// RndPcg は、math/rand/v2のPCGを利用してローカル乱数生成器を使うサンプルです。
//
// - PCG    : 速くて統計的に良い、だけど暗号的には弱い
// - ChaCha8: ChaCha20 の軽量版で、かなり強い（準暗号的）乱数を生成
//
// ただし、公式ドキュメントでは「セキュリティ用途には math/rand/v2 ではなく crypto/rand を使え」と書かれています。
//
// PCG も含めて擬似乱数生成器は「同じ seed（初期状態）＋同じ呼び出し順」なら同じ乱数列を再生成できます。
//
// # REFERENCES
//   - https://pkg.go.dev/math/rand/v2@go1.26.4#New
//   - https://pkg.go.dev/math/rand/v2@go1.26.4#NewPCG
func RndPcg() error {
	var (
		src1 = rand.NewPCG(123, 456)
		src2 = rand.NewPCG(123, 456)
		rnd1 = rand.New(src1)
		rnd2 = rand.New(src2)

		w = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	)
	defer w.Flush()

	fmt.Fprintf(w, "rnd1:\t%v\t%v\t%v\n", rnd1.IntN(100), rnd1.IntN(100), rnd1.IntN(100))
	fmt.Fprintf(w, "rnd2:\t%v\t%v\t%v\n", rnd2.IntN(100), rnd2.IntN(100), rnd2.IntN(100))

	return nil
}
