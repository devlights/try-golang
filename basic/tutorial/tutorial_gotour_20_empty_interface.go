package tutorial

import "fmt"

type (
	myIf interface {
	}

	myIfImpl struct {
	}

	data struct {
		value string
	}
)

func (d *data) String() string {
	return fmt.Sprintf("[data] value:%v", d.value)
}

func (m *myIfImpl) String() string {
	return "myIfImpl"
}

func GoTourEmptyInterface() error {
	// ------------------------------------------------------------
	// 空のインターフェース
	// Go言語において、メソッドを一つも持たないインターフェースを interface{} で表す.
	// このインターフェースを空のインターフェースと呼ぶ。
	//
	// Go言語では、この空のインターフェースは他の言語の Object 型に相当する.
	// つまり、 interface{} は任意の型の値を保持出来る.
	// ------------------------------------------------------------
	var (
		v interface{}
	)

	v = 100
	p(v)

	v = "helloworld"
	p(v)

	v = true
	p(v)

	v = &data{value: "data-value"}
	p(v)

	v = &myIfImpl{}
	p(v)

	var inf myIf = v
	p(inf)

	return nil
}

func p(v interface{}) {
	fmt.Printf("%v\n", v)
}
