// Option Pattern についてのサンプルです。
//
// #REFERENCES
//   - https://dev.to/c4r4x35/options-pattern-in-golang-10ph
package main

import (
	"fmt"
	"time"

	"github.com/devlights/try-golang/examples/singleapp/option_pattern/1/config"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	var (
		c *config.Config
	)

	c = config.New(
		"172.16.0.111",
		8888,
		config.WithRecvTimeout(30*time.Second),
		config.WithSendTimeout(5*time.Second),
	)
	fmt.Println(c)

	c = config.New(
		"localhost",
		12345,
	)
	fmt.Println(c)

	return nil
}
