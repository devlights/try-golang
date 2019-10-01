package tutorial

import "fmt"

func Functions() error {
	// ------------------------------------------------------------
	// 関数は、０個以上の引数を受け取ることができ、０個以上の戻り値を返すことが出来る
	// 関数は、予約語 func を指定して signature を定義する。
	// ------------------------------------------------------------
	oneParamVoidReturn(100)
	fmt.Println(oneParamOneReturn(100))
	fmt.Println(multiParamMultiReturn(100, 200))
	fmt.Println(multiParamMultiReturnWithReturnNames(100, 200))

	var (
		x, y = 100, 200
	)

	swap(&x, &y)
	fmt.Println(x, y)

	return nil
}

// 引数を一つ受け取り、戻り値なしの関数
func oneParamVoidReturn(x int) {
	fmt.Println("oneParamVoidReturn", x)
}

// 引数を一つ受け取り、一つの戻り値を返す関数
func oneParamOneReturn(x int) int {
	return x * x
}

// 複数の引数を受け取り、複数の戻り値を返す関数
func multiParamMultiReturn(x, y int) (int, int, int) {
	return x, y, x + y
}

// 複数の引数を受け取り、複数の戻り値を返す関数だが、戻り値の部分にも予め名前を指定しておくことが出来る
// この場合、予め戻り値は確保されているので、関数を終了する場合 return とのみ記載することが可能
func multiParamMultiReturnWithReturnNames(x, y int) (rx, ry, rsum int) {
	rx = x
	ry = y
	rsum = x + y

	return
}

// 指定された 2つの値 を入れ替え
func swap(x, y *int) {
	*x, *y = *y, *x
}
