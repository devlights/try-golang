package effectivego

import "fmt"

type (
	effectivego09st01 struct {
		x, y int
		b    bool
		s    string
	}
)

// AllocationWithNew -- Effective Go - Allocation with new の 内容についてのサンプルです。
func AllocationWithNew() error {
	/*
		https://golang.org/doc/effective_go.html#allocation_new

		- Goには、newとmakeというメモリ割り当て用の組み込み関数がある。
		- newは、メモリを割り当てるが、他の言語のnewとは違い、初期化は行われない。
		  - それぞれの型におけるゼロ値で埋められる。
		  - C#などでは、newするとコンストラクタが動作するが、Goにはそのような動きはしない。単純に割り当ててゼロ値で埋める。
		- newは、メモリを割り当てて、ゼロ値で埋めて、そのポインタを返してくれる。
		- Goの標準ライブラリの型はゼロ値でもうまく動くように作られている。(例: bytes.Buffer, sync.Mutexなど)
		- newを使うとゼロ値で埋められた値が返ってくるので、それに対して各プロパティを設定したりする
		  - だが、実際は composite literals を用いて、初期化してしまうことの方が結構多い。
	*/
	// 以下の２つは同じ状態となる.
	//   - d は ゼロ値で埋められた実体
	//   - p は ゼロ値で埋められた実体のポインタ
	var (
		d effectivego09st01
		p *effectivego09st01
	)

	p = new(effectivego09st01)

	fmt.Printf("d:%p (%v)\tp:%p (%v)\n", &d, d, p, *p)

	return nil
}
