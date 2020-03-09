package tutorial24

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type (
	// サンプル用のエラーを表す型 (1)
	goTour24Error struct {
		When time.Time
		What string
	}

	// サンプル用のエラーを表す型 (2)
	goTour24Error2 struct {
		goTour24Error
	}

	// サンプル用のエラーを表す型 (3)
	goTour24Error4 struct {
		message string
		inner   error
	}
)

// Impl: error.Error
func (e *goTour24Error2) Error() string {
	return fmt.Sprintf("[error2] %s", e.What)
}

// Impl: fmt.Stringer.String
func (e *goTour24Error2) String() string {
	return fmt.Sprintf("[stringer2] %v, %s", e.When, e.What)
}

// Impl: error.Error
func (e *goTour24Error) Error() string {
	return fmt.Sprintf("[error] %s", e.What)
}

// Impl: error.Error
func (e *goTour24Error4) Error() string {
	return e.message
}

// Unwrap 関数は、特定のインターフェースとして定義されていない。(errors.Unwrapの中でその場で定義して判定される）
// errors.Unwrap() は、対象のデータが Unwrap() error というシグネチャのメソッドを持っているかどうかを判断して処理する。
func (e *goTour24Error4) Unwrap() error {
	return e.inner
}

var (
	// エラー定義は、必ずエラー型を独自で定義する必要はなく、以下のように errors.New を利用して簡易生成することも出来る
	goTour24Error3 = errors.New("error 3")
	goTour24Error5 = errors.New("error 1.13 (1)")
)

// Error は、 Tour of Go - Errors (https://tour.golang.org/methods/19) の サンプルです。
func Error() error {

	aboutGoError()
	fmt.Println("----------------------------------------------------")
	aboutGo113Error()

	return nil
}

func aboutGo113Error() {
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
	// 必要があった。 1.13までは言語自体に共通的なやり方が無かった。
	//
	// Go 1.13 より、errorsパッケージに以下の関数が追加された.
	//    - (1) Unwrap()
	//    - (2) Is()
	//    - (3) As()
	// どれも、ヘルパー関数的な位置づけとなっている。
	// Unwrap()は、エラーを内包している場合に、この名前で関数定義しておくと errors.Unwrap() で自動的に呼び出される.
	// Is()は、エラーを値として定義している場合に使う. 型に Is(error) bool というメソッドを定義しておくと errors.Is() で自動的に呼び出される.
	// As()は、エラーを型として定義している場合に使う. 型に As(interface{}) bool というメソッドを定義しておくと errors.As() で自動的に呼び出される.
	//
	// また、fmt の書式化文字列に %w というものが追加された.
	// このフォーマット指示子で error をフォーマットしておくと
	// 自動で Unwrap 可能な状態にしてくれる. (4)
	//
	// ただし、pkg/errors や golang.org/x/xerrors に存在するスタックトレース付きの書式化文字列はサポートされていない。
	// (: %v とか : %w とか %+v とか %+w とか)
	// ------------------------------------------------------
	// 1.13 以降 (1)
	//   errors.Unwrap ができたので共通手順で内包しているエラーが取り出せるようになった。
	e4 := &goTour24Error4{
		message: "e4",
		inner:   goTour24Error5,
	}

	innerError := errors.Unwrap(e4)
	if innerError != nil {
		fmt.Printf("(1) Root: %v\tInner: %v\n", *e4, innerError)
	}

	// 1.13 以降 (2)
	//   前は innerError == goTour24Error5 としていた
	if errors.Is(innerError, goTour24Error5) {
		fmt.Printf("(2) innerError type: %T\n", innerError)
	}

	// 1.13 以降 (3)
	makeError := func() error {
		return &goTour24Error2{
			goTour24Error: goTour24Error{
				When: time.Now(),
				What: "error raised2",
			},
		}
	}

	e5 := &goTour24Error4{
		message: "e5",
		inner:   makeError(),
	}

	innerError2 := errors.Unwrap(e5)
	if innerError2 != nil {
		fmt.Printf("(3-1) Root: %v\tInner: %v\n", *e5, innerError2)
	}

	var ie *goTour24Error2
	if errors.As(innerError2, &ie) {
		fmt.Printf("(3-2) innerError type: %T\n", ie)
	}

	// (4) %w フォーマット指示子
	myError3 := fmt.Errorf("my error 3")
	myError4 := fmt.Errorf("my error 4 (%w)", myError3)

	var myError5 = myError4

	fmt.Println("(4-0)", myError5)

	// myError4 は、内部で %w を用いて myError3 を組み込んでいるので Is() で判定可能
	if errors.Is(myError5, myError3) {
		fmt.Println("(4-1) errors.Is(myError5, myError3) == true")
	}

	// 当然 myError4 を Is() で聞いても問題ない
	if errors.Is(myError5, myError4) {
		fmt.Println("(4-2) errors.Is(myError5, myError4) == true")
	}

	// %w で書式化している場合、自動的に Unwrap 可能な状態になっている
	myError6 := errors.Unwrap(myError5)
	fmt.Println("(4-3)", myError6)

	switch {
	case errors.Is(myError6, myError4):
		fmt.Println("(4-4) myError6.Is(myError4) == true")
	case errors.Is(myError6, myError3):
		fmt.Println("(4-4) myError6.Is(myError3) == true")
	}
}

func aboutGoError() {
	// ------------------------------------------------------------
	// Go言語のerror型
	//
	// Goのエラー処理の考え方はとてもシンプル。
	// 例外とかではなく、ちゃんとエラーを関数から返して呼び元が適切に判断して処理する。
	//
	// Go言語ではエラーの状態を error値 で表現する.
	// error型は fmt.Stringer インターフェースのように、これもインターフェースである.
	//
	// fmt.Stringerと同様に、fmtパッケージの関数は error インターフェースを
	// 意識して処理してくれる。 指定した値が error インターフェースを実装している場合
	// Error()を呼び出してくれる。
	// ちなみに、fmt.Stringerとerrorの両方のインターフェースを実装している場合
	// Error() が呼ばれる.
	//
	// errorsパッケージには、 error を簡易に生成することが出来る関数がある.
	//   - errors.New
	// また、fmtパッケージにも、 error を返してくれる関数がある。
	//   - fmt.Errorf
	//
	// Goでは、他の言語によくある try-catch の仕組みが意図的に存在しないため
	// 定義した関数の戻り値に error を追加して返すことが非常に多い。
	//
	// 呼び元は、関数を呼び出し、戻り値で error を受け取り、それが nil か否かを
	// 判定して、エラーが発生したかどうかを判定する。
	//
	// Goのエラー処理に関しては、元々搭載されているエラー機構が「ネストしたエラーの取り扱い」などに
	// 少し難があるため、 pkg/errors (https://github.com/pkg/errors) を利用している人も多い。
	// しかし、Go2に向けてエラー処理も使いやすいように変化しているので、今は過渡期と言えるかもしれない。
	// 実際、Go 1.13 で追加された errors.Unwrap や Is, As, %w などは、元々 pkg/errors に存在している機能である。
	//
	// Go2のプロポーサルの実装は、以下のパッケージで利用することが出来る。
	//   xerrors (golang.org/x/xerrors)
	//
	// エラー処理については、Go公式ブログに記事が上がっている
	//   - https://blog.golang.org/go1.13-errors
	//   - https://blog.golang.org/errors-are-values
	//   - https://blog.golang.org/error-handling-and-go
	//   - https://go.googlesource.com/proposal/+/master/design/29934-error-values.md
	//   - https://godoc.org/golang.org/x/xerrors
	//
	// また、以下の情報もとてもわかり易く書かれている。感謝。
	//   - https://qiita.com/sonatard/items/9c9faf79ac03c20f4ae1
	//   - https://qiita.com/sonatard/items/c9b985c2022cf5e438e9
	//   - https://qiita.com/sonatard/items/95c7a68eb1a378734b01
	//
	// ------------------------------------------------------------
	// よくあるエラー処理のやり方.
	s := "abc"
	_, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("数値変換に失敗 [%q][%v]\n", s, err)
	}

	// エラーを返す用の関数定義 (1)
	//   - 自前定義のエラー型
	run := func() error {
		return &goTour24Error{
			When: time.Now(),
			What: "error raised",
		}
	}

	if err = run(); err != nil {
		fmt.Println("(1)", err)
	}

	// エラーを返す用の関数定義 (2)
	//   - 自前定義のエラー型（ネスト版）
	run2 := func() error {
		return &goTour24Error2{
			goTour24Error: goTour24Error{
				When: time.Now(),
				What: "error raised2",
			},
		}
	}

	if err = run2(); err != nil {
		fmt.Println("(2-1)", err)

		// わざと interface{} にキャストし直して Println してみる
		if i, ok := err.(*goTour24Error2); ok {
			fmt.Println("(2-2)", i)
		}
	}

	// エラーを返す用の関数定義 (3)
	//   - errors.New() を利用
	run3 := func() error {
		return errors.New("errors.New")
	}

	if err = run3(); err != nil {
		fmt.Println("(3)", err)
	}

	// エラーを返す用の関数定義 (4)
	//   - fmt.Errorf() を利用
	run4 := func() error {
		return fmt.Errorf("%s", "fmt.Errorf")
	}

	if err = run4(); err != nil {
		fmt.Println("(4)", err)
	}

	// エラーを返す用の関数定義 (5)
	//   - 事前に定義しておいた errors.New を利用
	run5 := func() error {
		return goTour24Error3
	}

	if err = run5(); err != nil {
		if err == goTour24Error3 {
			fmt.Println("(5)", err)
		}
	}
}
