// Go の プログラムは、パッケージで構成される。
// 規約により、パッケージ名はimportパスの最後の要素の名前となる。
// プログラムは、必ず main パッケージから開始される。
// Go では、一つのディレクトリ内に一つのパッケージしか含めることが出来ない。
package tutorial

// 利用するパッケージのimportを記載する。
// 複数のimportがある場合は、以下のようにグループ化して記述することができる。
import (
	"fmt"
	"math"
	math2 "math" // alias設定
)

// GoTour は、 [A Tour of Go](http://bit.ly/2HsCMiG) の 要約.
func GoTour() error {
	// ------------------------------------------------------------
	// Hello World
	//   文字列を出力するには、 fmt パッケージの Println() などを使う
	// ------------------------------------------------------------
	fmt.Println("Hello World")

	// ------------------------------------------------------------
	// import したパッケージの利用
	//   import したパッケージは、その名前で利用することが出来る。
	//   (python などと同様)
	//   alias 定義した名前も利用できる。
	// ------------------------------------------------------------
	fmt.Println(math.Pi)
	fmt.Println(math2.Pi)

	// ------------------------------------------------------------
	// Goでは、最初の文字が大文字で始まる名前は外部に公開される。(public扱い)
	// 小文字で始まる名前は外部に公開されない。(private扱い)
	// 公開範囲は、パッケージ単位。なので、小文字の名前をつけた要素も同一パッケージ内
	// では、見ることが出来る。
	// ------------------------------------------------------------
	fmt.Println(packagePrivateFunc())

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

	// ------------------------------------------------------------
	// Go言語の基本型は以下
	//
	// - bool
	// - string
	// - int     (int8,   int16,  int32,  int64)
	// - uint    (uint8 , uint16, uint32, uint64, uintptr)
	// - byte    (uint8 の 別名)
	// - rune    (int32 の 別名)
	// - float   (float32, float64)
	// - complex (complex32, complex64)
	// ------------------------------------------------------------
	//noinspection GoVarAndConstTypeMayBeOmitted
	var (
		boolVal bool   = true
		maxInt  int64  = 1<<63 - 1
		maxUInt uint64 = 1<<64 - 1
		byteVal byte   = 42
		runeVal rune   = 'a'
	)

	//noinspection GoBoolExpressions
	fmt.Println(boolVal, maxInt, maxUInt, byteVal, runeVal)

	// ------------------------------------------------------------
	// 初期値
	// 変数に初期値を与えずに宣言すると、「ゼロ値」が設定される
	// ゼロ値は型によって異なる
	//
	// - 数値型は 0
	// - bool型は false
	// - string型は ""
	// ------------------------------------------------------------
	var (
		zeroInt    int
		zeroBool   bool
		zeroString string
	)

	fmt.Printf("%v\t%v\t%q\n", zeroInt, zeroBool, zeroString)

	return nil
}

// 小文字で始まっているのでこの関数は非公開関数 (パッケージプライベート)
func packagePrivateFunc() string {
	return "This is package-private function"
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
