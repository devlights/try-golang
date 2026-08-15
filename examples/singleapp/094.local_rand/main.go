package main

import (
	"fmt"
	// math/rand/v2 はシミュレーションやテスト向けの擬似乱数であり、暗号用途には適していない。
	"math/rand/v2"
)

func main() {
	var (
		rnd1 *rand.Rand
		rnd2 *rand.Rand
	)
	rnd1 = rand.New(rand.NewPCG(123, 345))
	rnd2 = rand.New(rand.NewPCG(123, 345))

	// 乱数生成器にPCG(Permuted Congruential Generator)を使うと、同じシート値を指定した場合は
	// 全く同じ乱数を得ることが出来る。開発時試験などで便利。
	// (注意) math/rand/v2 は、「真の乱数」ではなく、状態とシードから決まる決定的な擬似乱数であることに注意。
	//       本当の乱数を利用する場合は crypto/rand を利用すること。
	for range 5 {
		fmt.Println(rnd1.Int32(), rnd2.Int32())
	}
}
