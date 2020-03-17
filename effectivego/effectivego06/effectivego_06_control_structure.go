package effectivego06

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// ControlStructure -- Effective Go - Control structures の 内容についてのサンプルです。
func ControlStructure() error {
	/*
		https://golang.org/doc/effective_go.html#control-structures

		- Goにはループ制御構文として for のみが存在する。 while は存在しない。
		- Goのswitchは、他の言語よりフレキシブルに使える。分岐が多い場合は if より switch の方が見やすい。
		- if および switch には for のように初期化ステートメントを付与して評価することができる
		- Goのswitch は fall-through しない。なので、breakを書かなくても自動で落ちていかない
		- break および continue には、 追加でラベルを付与して、そこにジャンプさせることができる
		- select は、Go特有の制御構造。チャネルを制御する際に使う。
		- Goの制御構造には、他の言語 (CやC#やJavaなど）のように 条件部 を カッコで囲む必要はない
	*/

	// --------------------------------------------------------
	// if
	// --------------------------------------------------------
	ifExample()

	// -------------------------------------------------------------
	// for
	// -------------------------------------------------------------
	forExample()

	// -------------------------------------------------------------
	// switch
	// -------------------------------------------------------------
	switchExample()

	// -------------------------------------------------------------
	// select
	// -------------------------------------------------------------
	selectExample()

	return nil
}

func selectExample() {
	type (
		empty interface{}
	)
	done := make(chan empty)
	defer close(done)
	asyncF := func(done <-chan empty) <-chan empty {
		term := make(chan empty)

		go func() {
			defer close(term)

			log.Println("[asyncF] 非同期処理開始")

			select {
			case <-done:
				log.Println("[asyncF] 上位から終了命令がきたので終わり")
			case <-time.After(2 * time.Second):
				log.Println("[asyncF] 処理終わり")
			}
		}()

		return term
	}
	asyncDone := asyncF(done)
	log.Println("[main] 非同期処理 待ち合わせ開始")
	defer log.Println("[main] 非同期処理 待ち合わせ終了")
LOOP:
	for {
		select {
		case <-time.After(1 * time.Second):
			log.Println("[main] タイムアウト迎えたので非同期処理打ち切り")
			done <- new(empty)
		case <-asyncDone:
			log.Println("[main] 非同期処理が終わったとの連絡を受信")
			break LOOP
		default:
		}
	}
}

func switchExample() {
	// if のように利用
	i := rand.Intn(10)
	switch {
	case i == 0:
		fmt.Println("zero")
	case i <= 5:
		fmt.Println("less than 5")
	case i <= 8:
		fmt.Println("less then 8")
	default:
		fmt.Println("other")
	}
	// いつもの switch の形
	switch rand.Intn(5) {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}
	// 型判定のための switch (Type switch)
	// switch の 初期化部 で v.(type) のように 値.(type) と宣言すると型情報が取れるので
	// それで分岐させる
	//   https://golang.org/doc/effective_go.html#type_switch
	var v interface{}
	v = getSomeValue()
	switch t := v.(type) {
	case string:
		fmt.Println("type: string")
	case int:
		fmt.Println("type: int")
	case bool:
		fmt.Println("type: bool")
	default:
		fmt.Printf("other [%T]\n", t)
	}
}

func forExample() {
	// C言語ライク
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// for-each
	for index, value := range [3]int{1, 2, 3} {
		fmt.Printf("index[%d] value[%d]\n", index, value)
	}
	// mapの場合は key, value となる
	for key, value := range map[int]string{1: "one", 2: "two", 3: "three"} {
		fmt.Printf("key[%d] value[%v]\n", key, value)
	}
	// 文字列の場合、そのまま for-each すると rune がループ単位になる
	// 文字単位の場合は、必ずしも 1byte ずつループする訳ではないことに注意
	for pos, char := range "こんにちは" {
		fmt.Printf("pos[%d] char[%#U] code-point[0x%x]\n", pos, char, char)
	}
}

//noinspection GoBoolExpressions
func ifExample() {
	// 標準形
	i := 10
	if i <= 10 {
		fmt.Printf("i=%d\n", i)
	}

	// 初期化ステートメント付き
	f := "no_exists_file_path"
	if _, err := os.Stat(f); os.IsNotExist(err) {
		fmt.Printf("ファイル存在しない [%s]\n", f)
	}

	// -------------------------------------------------------------
	// Go で、 := を使った 暗黙的な変数宣言 を記述する場合
	// よく以下のように同じ errStat 変数を使いまわしたりする。
	// 暗黙的な変数宣言の再代入できる仕様については
	//   https://golang.org/doc/effective_go.html#redeclaration
	// を参照
	// -------------------------------------------------------------
	_, errStat := os.Stat(f)
	if errStat != nil {
		fmt.Println("errStat 1回目")
	}
	_, errStat = os.Stat(f)
	if errStat != nil {
		fmt.Println("errStat 2回目")
	}
}

func getSomeValue() interface{} {
	return false
}
