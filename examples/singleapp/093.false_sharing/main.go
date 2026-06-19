// False Sharingについてのサンプルです。
//
// False Sharingとは「別々の変数を触っているのに、たまたま同じキャッシュラインに乗っているせいで、キャッシュコヒーレンシが暴れて性能が落ちる現象」のこと。
// マルチコアCPUの環境で、一つの構造体の別々のフィールドを別々のスレッドで個別に更新する処理などを書いたりしている場合に良く発生します。
// 数回更新されるだけの処理なら問題ありませんが、例えば何万回、何千万回も更新されるフィールドの場合などは塵積で差が出てきます。
//
// キャッシュコヒーレンシとは「複数のキャッシュに載っている“同じメモリ位置のデータ”の内容を、常に矛盾なく保つための仕組み・性質」のことです。
// マルチコアCPUでは、各コアがそれぞれL1/L2キャッシュを持ち、同じアドレスのデータをローカルにキャッシュします。
// どれかのコアがそのデータを書き換えたときに、他コアのキャッシュ内容も古いままにならないように
// 「同じアドレスならどのコアから見ても同じ値になる」ことを保証するメカニズムがキャッシュコヒーレンシです。
// False Sharingが発生している場合は、これが何回も発生します。
//
// 多くのアーキテクチャでキャッシュラインサイズは 64 バイトです。
// なので、同じキャッシュラインに入っている値を、別々のスレッドで更新するとキャッシュコヒーレンシが発生します。
// 一つのスレッドが更新するたびに、同じキャッシュラインをコピーして保持しているスレッド全部に発生するため、性能が落ちていきます。
//
// 回避策はそれぞれ別のキャッシュラインに載るように適切にパディングしたりすることです。
//
// REFERENCES:
//   - https://abhisheklearn12.github.io/blogs/memory.html
//   - https://jp.xlsoft.com/documents/intel/compiler/17/cpp_17_win_lin/GUID-4B0549F1-045C-47B0-BEAD-872522D62863.html
//   - https://en.wikipedia.org/wiki/False_sharing
package main

import (
	"flag"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	Iterations    = 200000000 // 各カウンタをインクリメントする回数
	CacheLineSize = 64        //キャッシュラインのサイズ
)

// badCounters は、悪い例を表す構造体です。
//
// a と b が同じキャッシュラインに載る可能性が高い構造体です。
type badCounters struct {
	a atomic.Int64
	b atomic.Int64
}

// goodCounters は、良い例を表す構造体です。
//
// a と b の間にパディングを挟んで別キャッシュラインに分離しています。
type goodCounters struct {
	a atomic.Int64
	_ [CacheLineSize - 8]byte
	b atomic.Int64
}

func inc(v *atomic.Int64) {
	for range Iterations {
		v.Add(1)
	}
}

func (me *badCounters) initialize() {
	me.a.Store(0)
	me.b.Store(0)
}

func (me *badCounters) incA() {
	inc(&me.a)
}

func (me *badCounters) incB() {
	inc(&me.b)
}

func (me *goodCounters) initialize() {
	me.a.Store(0)
	me.b.Store(0)
}

func (me *goodCounters) incA() {
	inc(&me.a)
}

func (me *goodCounters) incB() {
	inc(&me.b)
}

func main() {
	var (
		numProc = flag.Int("proc", -1, "runtime.GOMAXPROCSに指定する値（-1の場合はデフォルト設定を使います)")
	)
	flag.Parse()

	if *numProc > 0 {
		// シングルスレッドの場合、キャッシュコヒーレンシの大量発生は
		// 原理上発生しないので、badとgoodの差はほぼ無くなる。
		runtime.GOMAXPROCS(*numProc)
	}

	var (
		bad   = new(badCounters)
		good  = new(goodCounters)
		wg    sync.WaitGroup
		start time.Time
	)
	bad.initialize()
	good.initialize()

	//
	// False Sharingが発生するパターン
	//
	start = time.Now()
	{
		wg.Go(bad.incA)
		wg.Go(bad.incB)
		wg.Wait()
	}
	fmt.Printf("False Sharing: %v ms\n", time.Since(start).Milliseconds())

	//
	// パディングを設置して別々のキャッシュラインに載るようにするパターン
	//
	start = time.Now()
	{
		wg.Go(good.incA)
		wg.Go(good.incB)
		wg.Wait()
	}
	fmt.Printf("Padding      : %v ms\n", time.Since(start).Milliseconds())
}
