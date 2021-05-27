package gotour22

import "fmt"

type (
	ifGoTour22 interface {
		ifGoTour221
		ifGoTour222
	}

	ifGoTour221 interface {
		func1() string
	}

	ifGoTour222 interface {
		func2() string
	}

	ifGoTour221Impl struct {
	}

	ifGoTour222Impl struct {
	}

	ifGoTour22BothImpl struct {
	}

	ifGoTour22NotImpl struct {
	}
)

func (i *ifGoTour22BothImpl) func2() string {
	return fmt.Sprintf("[func2] %T", i)
}

func (i *ifGoTour22BothImpl) func1() string {
	return fmt.Sprintf("[func1] %T", i)
}

func (i *ifGoTour222Impl) func2() string {
	return fmt.Sprintf("[func2] %T", i)
}

func (i *ifGoTour221Impl) func1() string {
	return fmt.Sprintf("[func1] %T", i)
}

// TypeSwitch は、 Tour of Go - Type switches (https://tour.golang.org/methods/16) の サンプルです。
func TypeSwitch() error {
	// ------------------------------------------------------------
	// Type Switch
	// Go言語には、Type Switch という機能がある.
	// これは、 Type Assertions を switch で定義して直列に使用できる機能.
	// つまり、そのインターフェースの中の具象型への変換に関して
	// 複数の候補が存在する場合に、通常のswitch文のようにcase分けして
	// 判定できるという機能。 if 文を並べるよりも見やすい。
	//
	// 以下の構文で記述する.
	// 		switch v := i.(type) {
	// 		case T:
	// 			インターフェース i の値がT型に変換できた場合
	// 		case S:
	// 			インターフェース i の値がS型に変換できた場合
	// 		default:
	// 			上記以外の場合
	// 		}
	//
	// switch のところの i.(type) という部分であるが、ここは
	// (type)の部分を望みの型で指定するのではなく、このまま (type) と
	// 書く。この (type) の部分が case の部分の型で置き換えられて
	// 評価される。
	// ------------------------------------------------------------
	var (
		sli1 []interface{}
		sli2 []interface{}
	)

	sli1 = make([]interface{}, 0)
	sli1 = append(sli1, 100)
	sli1 = append(sli1, "helloworld")
	sli1 = append(sli1, true)

	for _, v := range sli1 {
		switch d := v.(type) {
		case int:
			fmt.Printf("[int] %v\n", d)
		case string:
			fmt.Printf("[string] %v\n", d)
		case bool:
			fmt.Printf("[bool] %v\n", d)
		}
	}

	sli2 = make([]interface{}, 0)
	sli2 = append(sli2, &ifGoTour221Impl{})
	sli2 = append(sli2, &ifGoTour222Impl{})
	sli2 = append(sli2, &ifGoTour22BothImpl{})
	sli2 = append(sli2, &ifGoTour22NotImpl{})

	for _, v := range sli2 {
		switch d := v.(type) {
		case ifGoTour22:
			fmt.Println(d.func1(), d.func2())
		case ifGoTour221:
			fmt.Println(d.func1())
		case ifGoTour222:
			fmt.Println(d.func2())
		default:
			fmt.Printf("[それ以外] %T\n", d)
		}
	}

	return nil
}
