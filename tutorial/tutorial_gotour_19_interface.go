package tutorial

import "fmt"

type (
	// サンプル用のインターフェース
	sumer interface {
		sum() int
	}

	// サンプル用インターフェースを実装する型
	sumImpl struct {
		x, y int
	}
)

// コンストラクタ関数
func newSumImpl(x, y int) *sumImpl {
	return &sumImpl{
		x: x,
		y: y,
	}
}

// sumer インターフェースを実装
func (s *sumImpl) sum() int {
	if s == nil {
		fmt.Println("[sumImpl] <nil>")
		return 0
	}

	return s.x + s.y
}

// fmt.Stringer インターフェースを実装
func (s *sumImpl) String() string {
	if s == nil {
		fmt.Println("[sumImpl] <nil>")
		return fmt.Sprintf("x:%v\ty:%v", 0, 0)
	}

	return fmt.Sprintf("x:%v\ty:%v", s.x, s.y)
}

// Interface は、 Tour of Go - Interfaces (https://tour.golang.org/methods/9) の サンプルです。
//noinspection GoNilness
func Interface() error {
	// ------------------------------------------------------------
	// Go言語のインターフェース
	// Go言語におけるインターフェースは、メソッドのシグネチャの集まりを定義しているもの.
	// Go言語には、クラスの概念は無いが、型に対するメソッドの定義とインターフェースで
	// シンプルなポリモーフィズムを実現している。
	//
	// 他の言語のインターフェースの概念と大きく異なるのは、Go言語ではインターフェースを
	// 明示的に implements しているという記載は必要なく、そのインターフェースに定義
	// されているメソッドと同じシグネチャを定義しておけばインターフェースを実装していること
	// になる点である.
	//
	// 例えば、fmtパッケージには、Stringerというインターフェースが定義されており
	// Stringerインターフェースは、一つだけメソッドを定義している
	//     String() string
	// 自身で定義した型でfmt.Stringerインターフェースを実装したいと思った場合は
	// 単純に String() string というメソッドを定義すれば良い。
	//
	// ちなみに、Go言語では メソッド を一つだけ持つインターフェースは
	// xxxerという名前で定義するのが暗黙のようである.
	//
	// このようにダックタイピング的に利用できて、便利なGo言語のインターフェースだが
	// 一点注意点があって、それは nil の扱い方について。
	//
	// Go言語のインターフェースは、概念的に（実際のオブジェクト, 実際の型)　という
	// タプルを内部で持っているような構造になっている。
	//
	// 宣言しただけで、何も設定していないインターフェースの場合
	// そのインターフェースの値は nil である。 つまり内部の値も型も nil。
	// この場合、 インターフェース == nil は True となる。
	//
	// そこに、実際の型で、値がnilなデータをインターフェースに設定すると
	// インターフェースの内部データは、値がnilで、型が実オブジェクトの型という形になる。
	// この場合、インターフェース == nil は　False となる。
	//
	// つまり、インターフェース == nil が True と判定されるのは
	// インターフェース内部の実オブジェクトの値と型の両方が nil な場合のときだけである。
	//
	// なので、インターフェース型でデータを扱っている場合、インターフェースとしては nil では
	// 無いけれども、内部のデータが nil ということはあり得る。そのため、インターフェースを
	// 実装している具象型のメソッドでは、レシーバが nil の場合を考慮する必要がある。
	// (インターフェース自体は nil では無いので、メソッドの呼び出しは可能であるため）
	//
	// 他のオブジェクト指向言語（C#やJava）などを経験している人からすると
	// このような場合は、NullPointerExceptionなどが発生するだろうと思っている
	// ところで、Go言語では普通にメソッドが呼べてしまう場合があるので注意が必要。
	//
	// nil な インターフェースは、具体的な値も型も保持していないので
	// nil インターフェースのメソッドを呼び出すと、ランタイムエラーとなる.
	//
	// 参考: http://bit.ly/2LsltPP, http://bit.ly/2LutdB2
	// ------------------------------------------------------------
	var (
		ifValueAndTypeBothNotNil sumer
		ifValueAndTypeBothNil    sumer
		ifValueNilAndTypeNotNil  sumer
		implNotNil               *sumImpl
		implNil                  *sumImpl
	)

	// 普通に具象データを作成して、それをインターフェースに設定して扱う
	implNotNil = newSumImpl(10, 20)
	ifValueAndTypeBothNotNil = implNotNil

	printSum(ifValueAndTypeBothNotNil)
	printSum(newSumImpl(30, 40))

	// 宣言だけして、何も設定していないインターフェースは、インターフェースの値自体が nil
	ifVal := ifValueAndTypeBothNil
	fmt.Printf("[ifValueAndTypeBothNil] value:%v\ttype:%T\tis nil?:%v\n", ifVal, ifVal, ifVal == nil)

	// printSum() には nil が引数で渡る
	printSum(ifVal)

	// 具象データのポインタを宣言しただけの場合、そのポインタが示す先がないので、そのデータは当然 nil となっている
	// このデータをインターフェースに設定すると、インターフェースの値部分は nil だが、型が埋まるので
	// インターフェース自体は nil ではなくなる
	ifValueNilAndTypeNotNil = implNil
	ifVal = ifValueNilAndTypeNotNil
	fmt.Printf("[ifValueNilAndTypeNotNil] value:%v\ttype:%T\tis nil?:%v\n", ifVal, ifVal, ifVal == nil)

	// インターフェース自体は nil ではないので、普通にメソッドの呼び出しが行える
	// 呼び出された側は、レシーバが nil の状態でメソッド呼び出しされる
	// なので、メソッドを実装する場合は レシーバ が nil の場合を考慮する必要がある.
	printSum(ifVal)

	return nil
}

func printSum(v sumer) {
	fmt.Println(v)

	if v != nil {
		fmt.Printf("[sum()] %v\n", v.sum())
	}
}
