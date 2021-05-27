package gotour12

import "fmt"

// Defer は、 Tour of Go - Defer (https://tour.golang.org/flowcontrol/12) の サンプルです。
func Defer() error {
	// ------------------------------------------------------------
	// Go言語の defer について
	// defer は、 Go言語の特徴的な機能の一つ.
	// defer ステートメントは、deferに渡した関数の実行を呼び出し元の
	// 関数の終わりまで遅延させる機能。
	//
	// 他の言語でいうと、関数レベルで設置した try-finally みたいな感じ.
	// defer は、Go言語でプログラムを記述する際に頻発する機能で
	// 例えば、I/O 処理や非同期処理などでよく利用する
	//
	// defer には、匿名関数も指定することができる。
	// defer は、関数の呼び出しを要求するので匿名関数を利用する場合
	//   defer func(){}()
	// と記載する必要がある
	//
	// defer は、内部でスタックされており
	// 関数内で後で defer を呼び出した順から着火されていく. (LIFO)
	// つまり、関数内で最初に defer したものは、最後に実行される.
	//
	// 注意点として
	//　　　defer に指定した関数の呼び出し評価は遅延されるが
	//   関数の引数に指定された引数の値は、遅延せずにその場で評価される
	//
	//   defer 内では、関数の return 変数 の値を読み書きできる
	// というのがある.
	//
	// 参考：
	// https://golang.org/doc/effective_go.html#defer
	// https://blog.golang.org/defer-panic-and-recover
	// ------------------------------------------------------------
	defer func() {
		fmt.Println("defer - begin")
	}()

	func1()
	defer func2()
	func3()

	// defer に指定された Println(i) の i　の値は
	// 遅延評価されず、すぐに評価されるので、0 が出力される
	i := 0
	defer fmt.Println(i)
	i++

	// defer 内で その関数の return変数 にアクセスできる
	fmt.Println("deferReadWriteReturnValue: ", deferReadWriteReturnValue())

	defer func() {
		fmt.Println("defer - end")
	}()

	return nil
}

func deferReadWriteReturnValue() (i int) {
	i = 0
	defer func() {
		// defer 内で return変数 を加算している
		// なので、この関数の本来の処理上で return した
		// 時点では、値は 1 だが、ここで更に加算されて
		// 結果として 2 が返る.
		i++
	}()

	i++

	return
}

func func3() {
	fmt.Println("func3()")
}

func func2() {
	fmt.Println("func2()")
}

func func1() {
	defer func() {
		fmt.Println("defer - func1")
	}()

	fmt.Println("func1()")
}
