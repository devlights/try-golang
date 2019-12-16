package tutorial

import (
	"errors"
	"fmt"
	"time"
)

type (
	goTour24Error struct {
		When time.Time
		What string
	}

	goTour24Error2 struct {
		goTour24Error
	}

	myError113After01 struct {
		message string
		inner   error
	}
)

func (e *goTour24Error2) Error() string {
	return fmt.Sprintf("[error2] %s", e.What)
}

func (e *goTour24Error2) String() string {
	return fmt.Sprintf("[stringer2] %v, %s", e.When, e.What)
}

func (e *goTour24Error) Error() string {
	return fmt.Sprintf("[error] %s", e.What)
}

func (e *myError113After01) Error() string {
	return e.message
}

func (e *myError113After01) Unwrap() error {
	return e.inner
}

// Error は、 Tour of Go - Errors (https://tour.golang.org/methods/19) の サンプルです。
func Error() error {
	// ------------------------------------------------------------
	// Go言語のerror型
	// Go言語ではエラーの状態を error値 で表現する.
	// error型は fmt.Stringer インターフェースのように、これもインターフェースである.
	//
	// fmt.Stringerと同様に、fmtパッケージの関数は error インターフェースを
	// 意識して処理してくれる。 指定した値が error インターフェースを実装している場合
	// Error()を呼び出してくれる。
	// ちなみに、fmt.Stringerとerrorの両方のインターフェースを実装している場合
	// Error() が呼ばれる.
	//
	// errorsパッケージには、 error に関する便利な関数がいろいろある.
	//   - errors.New
	// また、fmtパッケージにも、 error を返してくれる関数がある。
	//   - fmt.Errorf
	//
	// エラー処理については、Go公式ブログに記事が上がっている
	//   - https://blog.golang.org/go1.13-errors
	//   - https://blog.golang.org/errors-are-values
	//   - https://blog.golang.org/error-handling-and-go
	// ------------------------------------------------------------
	run := func() error {
		return &goTour24Error{
			When: time.Now(),
			What: "error raised",
		}
	}

	run2 := func() error {
		return &goTour24Error2{
			goTour24Error: goTour24Error{
				When: time.Now(),
				What: "error raised2",
			},
		}
	}

	run3 := func() error {
		return errors.New("errors.New")
	}

	run4 := func() error {
		return fmt.Errorf("%s\n", "fmt.Errorf")
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	if err := run2(); err != nil {
		fmt.Println(err)

		// わざと interface{} にキャストし直して Println してみる
		i := err.(interface{})
		fmt.Println(i)
	}

	if err := run3(); err != nil {
		fmt.Println(err)
	}

	if err := run4(); err != nil {
		fmt.Println(err)
	}

	// ------------------------------------------------------
	// Go 1.13 より前と 1.13 以降ではエラー処理のセオリーがちょっと異なる
	//   - https://blog.golang.org/go1.13-errors
	//
	// Go 1.13 より前
	//   - (1) どのエラーなのかを判定するためにパッケージ変数でエラーをしたりする
	//     - errors.New を利用して変数定義
	//   - (2) 明示的に 独自のエラー型 を定義して、型検証して判定したりする
	//
	// 上記のどちらの場合でも、エラーを多段階で内包しているような場合
	// 元の情報を取り出すという過程が必要となる場合がある。
	// しかし、(1)の場合は既に値がテキストとなっているので、なくなってしまっっている.
	// また、(2)の場合も、内部に 内包するエラー　を 保持して、自分で参照したりする
	// 必要があった。
	//
	// Go 1.13 より、errorsパッケージに以下の関数が追加された.
	//    - (1) Unwrap()
	//    - (2) Is()
	//    - (3) As()
	// どれも、ヘルパー関数的な位置づけとなっている。
	// Is()は、エラーを値として定義している場合に使う.
	// As()は、エラーを型として定義している場合に使う.
	//
	// また、fmt の書式化文字列に %w というものが追加された.
	// このフォーマット指示子で error をフォーマットしておくと
	// 自動で Unwrap 可能な状態にしてくれる. (4)
	// ------------------------------------------------------
	// 1.13 より前 (1)
	myError1 := errors.New("my error before 1.13 (1)")
	myError1Func := func() error {
		return myError1
	}

	if err := myError1Func(); err == myError1 {
		fmt.Println("my error1 occured")
	}

	// 1.13 より前 (2)
	myError2Func := func() error {
		return run2()
	}

	if err := myError2Func(); err != nil {

		switch err.(type) {
		case *goTour24Error:
			fmt.Println("goTour24Error has raised")
		case *goTour24Error2:
			fmt.Println("goTour24Error2 has raised")
		default:
			fmt.Println("others...")
		}
	}

	// 1.13 以降 (1)
	e113after01 := &myError113After01{
		message: "myError113After01",
		inner:   myError1Func(),
	}

	innnerError := errors.Unwrap(e113after01)
	if innnerError != nil {
		fmt.Printf("Root: %v\n\t%v\n", *e113after01, innnerError)
	}

	// 1.13 以降 (2)
	if errors.Is(innnerError, myError1) {
		fmt.Printf("(2-1) innerError type: %T\n", innnerError)
	}

	// 1.13 以降 (2)
	e113after02 := &myError113After01{
		message: "myError113After02",
		inner:   myError2Func(),
	}

	innerError2 := errors.Unwrap(e113after02)
	if innerError2 != nil {
		fmt.Printf("Root: %v\n\t%v\n", *e113after02, innerError2)
	}

	var ie *goTour24Error2
	if errors.As(innerError2, &ie) {
		fmt.Printf("(2-2) innerError type: %T\n", ie)
	}

	// (4) %w フォーマット指示子
	myError3 := fmt.Errorf("my error 3")
	myError4 := fmt.Errorf("my error 4 (%w)", myError3)

	var myError5 error = myError4

	fmt.Println(myError5)

	// myError4 は、内部で %w を用いて myError3 を組み込んでいるので Is() で判定可能
	if errors.Is(myError5, myError3) {
		fmt.Println("errors.Is(myError5, myError3) == true")
	}

	// 当然 myError4 を Is() で聞いても問題ない
	if errors.Is(myError5, myError4) {
		fmt.Println("errors.Is(myError5, myError4) == true")
	}

	// %w で書式化している場合、自動的に Unwrap 可能な状態になっている
	myError6 := errors.Unwrap(myError5)
	fmt.Println(myError6)

	switch {
	case errors.Is(myError6, myError4):
		fmt.Println("myError6.Is(myError4) == true")
	case errors.Is(myError6, myError3):
		fmt.Println("myError6.Is(myError3) == true")
	}

	return nil
}
