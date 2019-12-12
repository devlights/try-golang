package effectivego

// Effective Go - Allocation with new の 内容についてのサンプルです。
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
		  - が、実際は composite literals を用いて、初期化してしまうことの方が結構多い。
	*/

	return nil
}
