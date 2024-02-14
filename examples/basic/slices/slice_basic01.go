package slices

import "fmt"

// Basic01 -- スライスについてのサンプル
func Basic01() error {
	// int のスライスを宣言
	// GO言語では、サイズ指定しているものは「配列」
	// スライスは配列への参照を持つデータ構造のこと。
	//
	// スライスは型表記させると []int のようになる
	// 配列は型表記させると [3]int のようになる

	// 以下はスライスの例
	var l1 []int
	fmt.Printf("l1: %T\n", l1)

	for i := 0; i < 5; i++ {
		l1 = append(l1, i)
	}

	for _, item := range l1 {
		fmt.Println(item)
	}

	fmt.Printf("\n\n")

	// 以下は配列の例
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1: %T\n", a1)

	for _, item := range a1 {
		fmt.Println(item)
	}

	// スライスは参照なので別の変数に代入すると
	// 両方とも同じ参照を示す。(Java や C# や Python と同じ)
	l2 := l1
	//noinspection GoNilness
	fmt.Printf("l1: %d\nl2: %d\n", &l1[0], &l2[0])

	// 配列は実態なので別の変数に代入すると
	// 値のコピーが発生する
	a2 := a1
	fmt.Printf("a1: %d\na2: %d\n", &a1[0], &a2[0])

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: slice_basic01

	   [Name] "slice_basic01"
	   l1: []int
	   0
	   1
	   2
	   3
	   4


	   a1: [3]int
	   1
	   2
	   3
	   l1: 824633901952
	   l2: 824633901952
	   a1: 824633909968
	   a2: 824633910016


	   [Elapsed] 37.17µs
	*/

}
