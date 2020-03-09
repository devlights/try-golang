package shadowing

import (
	"fmt"

	"github.com/devlights/try-golang/output"
)

// Basic は、変数のshadowingについてのサンプルです。
//
// REFERENCES::
//   - https://devlights.hatenablog.com/entry/2020/03/04/183433
func Basic() error {
	// --------------------------------------------------------------
	// 変数のshadowingが発生すると一見ちゃんとしているように見えて
	// 後で値を見たときに、予想外の状態になっていることが多い。
	//
	// 以下では、変数 hoge を最初にポインタで宣言しているが
	// 初期化をしていないので、 すぐ下のifにて := 代入の時点で
	// ifのスコープの変数が定義されてしまい、ifのスコープを抜けた後の
	// 値を見ると nil のままになっているというもの。
	// --------------------------------------------------------------
	output.Stdoutl("[shadowing   ] --------------------------------------")
	shadowing()
	output.Stdoutl("[no shadowing] --------------------------------------")
	noShadowing()

	return nil
}

func do(v string) (*string, error) {
	return &v, nil
}

func shadowing() {

	var hoge *string
	if true {
		// ここで shadowing が発生している
		// GoLand の Shadowing variable をONにしていると検知してくれる
		//noinspection GoShadowedVar
		hoge, err := do("word")
		if err != nil {
			return
		}

		// if の スコープの中では値がちゃんと見えている (wordって表示される)
		fmt.Printf("CHECKPOINT: %v(%v)\n", *hoge, hoge)
	} else {
		hoge = nil
	}

	// ここは if のスコープを抜けているので値がnilになる
	//noinspection GoNilness
	if hoge != nil {
		fmt.Printf("RESULT    : %v(%v)\n", *hoge, hoge)
	} else {
		fmt.Printf("RESULT    : nil\n")
	}
}

func noShadowing() {

	var hoge *string
	if true {
		result, err := do("word")
		if err != nil {
			return
		}

		hoge = result

		fmt.Printf("CHECKPOINT: %v(%v)\n", *hoge, hoge)
	} else {
		hoge = nil
	}

	if hoge != nil {
		fmt.Printf("RESULT    : %v(%v)\n", *hoge, hoge)
	}
}
