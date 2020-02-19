package effectivego

import "fmt"

// Effective Go - Two Dimensional Slices の 内容についてのサンプルです。
func TwoDimensionalSlices() error {
	/*
		https://golang.org/doc/effective_go.html#two_dimensional_slices

		- Two Dimensional Slicesとは、２次元スライスのこと
		- ２次元の配列の場合、変数宣言時に全次元がゼロ値が埋まる
		- ２次元のスライスの場合、最初のmake()で生成されるのは１次元目のみ。２次元目はまだnilのまま。
		  - 自分でその都度割り当てていく
	*/
	separator := func() {
		fmt.Println("----------------------")
	}

	dumpArray := func(a [2][2]int) {
		for i, v1 := range a {
			for j, v2 := range v1 {
				fmt.Printf("[%d][%d] %v1\n", i, j, v2)
			}
		}
	}

	dumpSlice := func(s [][]rune) {
		for i, v1 := range s {
			if v1 == nil {
				fmt.Printf("[%d][x] is nil\n", i)
			}

			for j, v2 := range v1 {
				fmt.Printf("[%d][%d] %v1\n", i, j, v2)
			}

			fmt.Println(string(s[i]))
		}
	}

	// ------------------------------------------------
	// ２次元配列の場合
	// ------------------------------------------------
	array2D := [2][2]int{}
	dumpArray(array2D)
	separator()

	array2D[0][1] = 100
	array2D[1][0] = 999

	dumpArray(array2D)
	separator()

	// ------------------------------------------------
	// ２次元スライスの場合
	// ------------------------------------------------
	slice2D := make([][]rune, 2)
	dumpSlice(slice2D)
	separator()

	// スライスの場合は配列と違い２次元目はまだ生成されていないので個別で割り当てる必要がある
	for i, v := range []string{"hello", "世界"} {
		slice2D[i] = []rune(v)
	}

	dumpSlice(slice2D)
	separator()

	return nil
}
