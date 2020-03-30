package effectivego29

import (
	"sync"

	"github.com/devlights/gomy/output"
)

const (
	Poison = "die"
)

// PanicRecover -- Effective Go - Recover の 内容についてのサンプルです。
func PanicRecover() error {
	/*
		https://golang.org/doc/effective_go.html#recover

		- panic() が呼ばれるとプログラムは強制終了する
		- 基本的に panic() は利用するべきではない。ちゃんとエラーを返すべき。
		  - panic() を利用するべきなのは、アプリケーションが再起不能と判定できるとき
		  - 利用されるべきシーンとしては、初期化処理など。初期化処理の場合は、ちゃんと成功しないとその先処理を進めるべきではないため。
		- recover() を使うことで panic が発生した goroutine の制御を回復して、最後のチャンスを貰える
		  - panic は当該 goroutine の実行を強制終了してスタックを巻き戻す。recover はこの伝播をストップしてくれる。
		  - panic が発生した際にその goroutine にて recover が存在しない場合はスタックが巻き戻ってランタイムへと伝播し、プログラムが落ちる.
		- recover() は defer の中でのみ利用できる
		- recover() は goroutine 単位で仕込む
	*/
	// それぞれのデータを非同期処理し、途中のデータで panic するようにしておく
	start("hello", Poison, "world").Wait()

	return nil
}

func start(values ...string) *sync.WaitGroup {
	wg := sync.WaitGroup{}
	wg.Add(len(values))

	for _, v := range values {
		output.Stdoutl("[proc][start]", v)
		go procWithRecover(&wg, v)
	}

	return &wg
}

func procWithRecover(wg *sync.WaitGroup, data string) {
	defer wg.Done()
	defer func() {
		// recovert は defer の中で利用する
		if err := recover(); err != nil {
			output.Stdoutl("[recover]", err)
		}
	}()

	// 以下の関数は特定の条件で panic する
	proc(data)
	output.Stdoutl("[proc][end  ]", data)
}

func proc(data string) {
	if data == Poison {
		output.Stdoutl("[panic]", data)
		panic(data)
	}

	output.Stdoutl("[proc]", data)
}
