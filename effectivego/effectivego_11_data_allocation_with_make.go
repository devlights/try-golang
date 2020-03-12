package effectivego

import (
	"fmt"
	"time"
)

// AllocationWithMake -- Effective Go - Allocation with new の 内容についてのサンプルです。
func AllocationWithMake() error {
	/*
		https://golang.org/doc/effective_go.html#allocation_make

		- make() は、slice, map, channelの生成のみに用いる
		- make() は、値を生成して「初期化」して返す。（new()はゼロ値で埋める)
		- make() は、ポインタを返さない。(new()はポインタを返す)

		var p *[]int = new([]int)			// new()で生成しているのでスライスの構造体のゼロ値を返す。つまり nil.
		var v  []int = make([]int, 0, 0)	// make() で生成しているので実体を返す. len=0, cap=0
	*/
	// slice
	//   new()とmake()の違い
	//     現実的にスライスをnew()で生成することにメリットはないので、ほぼ使わない
	var (
		p  = new([]int)
		v  = make([]int, 0, 0)
		v2 = []int{1, 2}
	)

	fmt.Printf("p:%v nil?(%v)\tv:%v\tv2:%v\n", *p, *p == nil, v, v2)

	// map
	var (
		m  = make(map[string]int)
		m2 = map[string]int{
			"hello": 1,
			"world": 2,
		}
	)

	fmt.Printf("m:%v\tm2:%v\n", m, m2)

	// channel
	var (
		c = make(chan int)
	)

	go func() {
		defer close(c)
		time.Sleep(1 * time.Second)
		c <- 10
	}()

	fmt.Printf("c:%v\n", <-c)

	return nil
}
