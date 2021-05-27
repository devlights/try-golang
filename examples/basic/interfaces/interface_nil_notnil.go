package interfaces

import (
	"strconv"

	"github.com/devlights/gomy/output"
)

// NilOrNotNil -- interface{} が nil になったり not-nil になったりする場合のサンプルです
//
// REFERENCES:
//   - https://medium.com/@shivi28/go-when-nil-nil-returns-true-a8a014abeffb
func NilOrNotNil() error {
	//
	// Go の interface は 型と値を内部でもつ
	// 両方が nil の場合に nil となる
	// なので、型がnot-nilで、値がnilの場合は nil とならない
	//

	// ポインタ型のゼロ値は nil
	// 型と値の表示でいうと、 <*int, nil> となる
	var i *int
	// それを interface{} に代入すると型は *int で、値が nil となる
	// つまり、<*int, nil>
	var o interface{} = i
	// interface のゼロ値は <nil, nil> となる
	var o2 interface{}

	output.Stdoutl("[i  == nil]", i == nil)  // true
	output.Stdoutl("[o  == nil]", o == nil)  // false (値はnilだけど型がnot-nil)
	output.Stdoutl("[o2 == nil]", o2 == nil) // true  (値も型もnil)

	output.Stdoutl("[i == o ]", i == o)  // true  (どちらも型が *int で、値が nil)
	output.Stdoutl("[i == o2]", i == o2) // false (片方は *int(nil)で、片方は nil(nil))

	nilnil := returnNilNil()
	notnilnil := returnNotnilNil()

	output.Stdoutl("[returnNilINil()   == nil]", nilnil == nil)
	output.Stdoutl("[returnNotnilNil() == nil]", notnilnil == nil)

	var e *myE
	output.Stdoutl("[nilnil    == e]", nilnil == e)
	output.Stdoutl("[notnilnil == e]", notnilnil == e)

	// 以下は panic する (<nil, nil>のため)
	// output.Stdoutl("[nilnil.Error()]", nilnil.Error())

	// 以下の呼び出しは panic しない。(<*myE, nil>のため)
	// ただし、値は<nil>なのでメソッドのレシーバーには<nil>が渡されて呼び出される
	output.Stdoutl("[notnilnil.Error()]", notnilnil.Error())
	output.Stdoutl("[e.Error()]", e.Error())

	tmp := myE(100)
	e = &tmp

	output.Stdoutl("[e.Error()]", e.Error())

	return nil
}

type myE int

func (e *myE) Error() string {
	if e == nil {
		return "<nil>"
	}

	return strconv.Itoa(int(*e))
}

func returnNilNil() error {
	// 返す値は <nil, nil>
	return nil
}

func returnNotnilNil() error {
	var e *myE

	// 返す値は <*myE, nil>
	return e
}
