package closure

import "fmt"

// Basic -- クロージャのサンプルです
func Basic() error {
	var (
		closure = func(i int) func() int {
			v := i
			return func() int {
				defer func() { v++ }()
				return v
			}
		}
		fn = closure(1)
	)

	for i := 0; i < 5; i++ {
		fmt.Println(fn())
	}
	
	return nil
}
