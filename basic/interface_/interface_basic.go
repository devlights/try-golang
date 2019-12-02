package interface_

import "fmt"

// サンプル用のインターフェース
type myInterface interface {
	// サンプル用のインターフェースメソッド
	myMethod() int
}

// サンプル用の構造体1
type mySt01 struct {
}

// サンプル用の構造体2
type mySt02 struct {
}

// impl: myInterface::myMethod
func (*mySt01) myMethod() int {
	return 10
}

// impl: myInterface::myMethod
func (*mySt02) myMethod() int {
	return 20
}

// Basic は、Goのインターフェースの基本に関するサンプルです.
// see also: tutorial/tutorial_gotour_19_interface.go
func Basic() error {
	var list []myInterface

	list = append(list, &mySt01{})
	list = append(list, &mySt02{})

	for i, item := range list {
		fmt.Printf("[%d]: [%d](%T)\n", i, item.myMethod(), item)
	}

	return nil
}
