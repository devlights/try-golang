package interface_

import "fmt"

// サンプル用のインターフェース
type MyInterface interface {
	// サンプル用のインターフェースメソッド
	MyMethod() int
}

// サンプル用の構造体1
type MySt01 struct {
}

// サンプル用の構造体2
type MySt02 struct {
}

// impl: MyInterface::MyMethod
func (*MySt01) MyMethod() int {
	return 10
}

// impl: MyInterface::MyMethod
func (*MySt02) MyMethod() int {
	return 20
}

// interface に関するサンプル
func Interface01() error {
	var list []MyInterface

	list = append(list, &MySt01{})
	list = append(list, &MySt02{})

	for i, item := range list {
		fmt.Printf("[%d]: [%d](%T)\n", i, item.MyMethod(), item)
	}

	return nil
}
