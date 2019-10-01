package tutorial

import (
	"fmt"
	"strings"
)

//noinspection GoUnusedType
type (
	ifGoTour20 interface {
		Value() string
	}

	ifGoTour20Impl struct {
		V string
	}

	ifGoTour20NotImpl struct {
		V string
	}
)

// impl fmt.Stringer
func (i *ifGoTour20Impl) String() string {
	return i.Value()
}

// impl tutorial.ifGoTour20
func (i *ifGoTour20Impl) Value() string {
	return strings.ToUpper(i.V)
}

func GoTourTypeAssertion() error {
	// ------------------------------------------------------------
	// Go言語の型検証 (Type Assertions)
	// 他の言語でもある型検証の仕組みがGo言語にも当然ある.
	// インターフェースに設定されている値の型を具象型または別のインターフェース
	// 型に変換する際に利用する.
	//
	// 基本的な構文としては以下のようになる
	//   t := i.(T)
	// インターフェースに設定されている値がT型に変換可能な場合はtが返る.
	//
	// この場合、変換できない場合にランタイムエラーとなる.
	// 変換可能かどうかがはっきりしていない場合は、マップにキーが存在するかを
	// 確認するのと同じような形で以下のように書く。
	//   t, ok := i.(T)
	// 変換できた場合は ok に true が設定される.
	//
	// 注意点として
	// 		type I interface {
	// 			V() string
	// 		}
	//
	// 		type Impl struct {
	// 		}
	//
	// 		func (i *Impl) V() string {
	// 			return "helloworld"
	//		}
	// という定義になっている状態でインターフェース I に具象型 Impl が
	// 設定されているとする。この場合、IからImplに変換する場合には
	// 		i.(Impl)
	// とするとコンパイルエラーとなる。具象型 Impl で実装している
	// インターフェースのメソッドは、ポインタレシーバーを受け取る様になっている
	// ため、正しく変換するには、以下のように書く。
	// 		i.(*Impl)
	// ------------------------------------------------------------
	var (
		i interface{} = "helloworld"
	)

	// 文字列型に変換
	if t, ok := i.(string); ok {
		fmt.Printf("[i.(string)] %T\t%v\n", t, t)
	}

	// 数値に変換
	// 変換できないので、ok は false になる
	if t, ok := i.(int); ok {
		fmt.Printf("[i.(int)] %T\t%v\n", t, t)
	}

	// 独自で定義したインターフェースの場合でも同じ
	var (
		impl            = &ifGoTour20Impl{V: "helloworld"}
		i2   ifGoTour20 = impl
	)

	// 具象型へ変換
	if t, ok := i2.(*ifGoTour20Impl); ok {
		fmt.Printf("[i2.(*ifGoTour20Impl)] %T\t%v\n", t, t)
	}

	// インターフェースを実装していない型に変換しようとすると
	// impossible type assertion: *ifGoTour20NotImpl does not implement
	// ifGoTour20 (missing Value method)
	// とコンパイルエラーとなる
	//if t, ok := i2.(*ifGoTour20NotImpl); ok {}

	// 変換元にインターフェース以外を設定すると
	// invalid type assertion: impl.(ifGoTour20) (non-interface type *ifGoTour20Impl on left)
	// とコンパイルエラーとなる.
	//if t, ok := impl.(ifGoTour20); ok {}

	// 別のインターフェースへの変換も出来る
	if t, ok := i2.(fmt.Stringer); ok {
		fmt.Printf("[i2.(fmt.Stringer)] %T\t%v\n", t, t)
	}

	// interface{} への変換は当然成功する
	if t, ok := i2.(interface{}); ok {
		fmt.Printf("[i2.(interface{})] %T\t%v\n", t, t)
	}

	return nil
}
