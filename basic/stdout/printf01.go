package stdout

import "fmt"

func Printf01() error {
	message := "Hello Golang!!"
	fmt.Printf("%s", message)

	return nil
}
