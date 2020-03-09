package closure

import "fmt"

// クロージャのサンプルです
//noinspection GoNameStartsWithPackageName
func Closure01() error {
	c1 := myClosure1()

	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	c2 := myClosure1()

	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())

	return nil
}

func myClosure1() func() int {
	var val = 0
	return func() int {
		val++
		return val
	}
}
