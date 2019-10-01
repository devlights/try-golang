package tutorial

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// Switch は、 Tour of Go - Switch (https://tour.golang.org/flowcontrol/9) の サンプルです。
func Switch() error {
	// ------------------------------------------------------------
	// Go言語の switch は、他の言語に比べてとても高機能で使いやすい.
	// 以下の特徴を持つ.
	//
	// - 自動で fall-through しないので、breakを書く必要がない
	// - case に指定するのは定数である必要がなく、関係する値も整数である必要がない
	//   - case f(): と書くこともできる
	// - if と 同様に 条件部に ステートメント も指定できる
	//   - switch xxx; xxx {} のように書ける
	// - 条件部を 指定しない switch　は、 switch true {} と書いてるのと同じ
	//   - if-then-else を長くかくより、 switch {} を使うほうがシンプルに表現できる.
	// ------------------------------------------------------------
	rand.Seed(time.Now().UnixNano())
	var (
		i = rand.Intn(6)
		f = func() int {
			return 3
		}
	)

	// 標準的な使い方
	switch myOs := runtime.GOOS; myOs {
	case "windows":
		fmt.Println("win")
	case "darwin":
		fmt.Println("mac")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("それ以外")
	}

	// case の部分に関数呼び出しも指定出来る
	switch i {
	case 1:
		fmt.Println("one")
	case f():
		fmt.Println("f()")
	default:
		fmt.Println("other")
	}

	// 条件部を省略
	switch {
	case i < f():
		fmt.Println("< ", f())
	case i > f():
		fmt.Println("> ", f())
	}

	return nil
}
