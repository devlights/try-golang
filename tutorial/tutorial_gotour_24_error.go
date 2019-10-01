package tutorial

import (
	"errors"
	"fmt"
	"time"
)

type (
	GoTour24Error struct {
		When time.Time
		What string
	}

	GoTour24Error2 struct {
		GoTour24Error
	}
)

func (e *GoTour24Error2) Error() string {
	return fmt.Sprintf("[error2] %s", e.What)
}

func (e *GoTour24Error2) String() string {
	return fmt.Sprintf("[stringer2] %v, %s", e.When, e.What)
}

func (e *GoTour24Error) Error() string {
	return fmt.Sprintf("[error] %s", e.What)
}

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
	// ------------------------------------------------------------
	run := func() error {
		return &GoTour24Error{
			When: time.Now(),
			What: "error raised",
		}
	}

	run2 := func() error {
		return &GoTour24Error2{
			GoTour24Error: GoTour24Error{
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

	return nil
}
