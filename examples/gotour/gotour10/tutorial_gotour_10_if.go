package gotour10

import (
	"fmt"
	"math/rand"
	"time"
)

// If は、 Tour of Go - If (https://tour.golang.org/flowcontrol/5) の サンプルです。
func If() error {
	// ------------------------------------------------------------
	// Go言語の if は、 他の言語の if の書き方とほぼ同様
	// C# などのように、 if () {} と書く必要はなくて、 if {} とカッコなしで書く
	// Go言語では、通常の if の書き方に加えて、ステートメント付きの if も書ける
	// これは、 if の条件部で実行したステートメントの値をその場で判定して利用できる
	// 以下のように書く.
	//
	// if statement; condition {
	// }
	//
	// このパターンは、Go言語においてエラー判定などで頻発する
	//
	// if err := xxxx(); err != nil {
	// }
	// ------------------------------------------------------------
	var (
		rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
		i   = rnd.Intn(100)
	)

	// 通常の if
	if i < 50 {
		fmt.Println("< 50")
	} else if i < 30 {
		fmt.Println("< 30")
	} else {
		fmt.Println("> 50")
	}

	// ステートメント付きの if
	// ifの条件部で宣言した変数は if ブロックの中で使えるし
	// elseブロックの中でも見える
	if i2 := rnd.Intn(100); i2 < 80 {
		fmt.Println("< 80")
	} else {
		fmt.Println("> 80")
	}

	return nil
}
