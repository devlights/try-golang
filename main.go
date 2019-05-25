package main

import (
	"bufio"
	"fmt"
	"github.com/devlights/try-golang/advanced/async"
	"github.com/devlights/try-golang/advanced/reflection"
	"github.com/devlights/try-golang/basic/helloworld"
	"github.com/devlights/try-golang/basic/map_"
	"github.com/devlights/try-golang/basic/scope"
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"log"
	"os"
	"strings"
)

func main() {
	mapping := makeMappings()

	fmt.Print("ENTER EXAMPLE NAME: ")

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		// 実行サンプル名取得
		example := stdinScanner.Text()
		if strings.ToLower(example) == "quit" {
			// 終了
			break
		}

		// サンプル実行
		if v, ok := mapping[example]; ok {
			if err := v(); err != nil {
				log.Fatal(err)
			}

			fmt.Print("\n\n")
		}

		fmt.Print("ENTER EXAMPLE NAME: ")
	}

	fmt.Println("END")
}

func makeMappings() map[string]func() error {
	mapping := make(map[string]func() error)

	mapping["helloworld"] = helloworld.HelloWorld
	mapping["printf01"] = stdout.Printf01
	mapping["printf02"] = stdout.Printf02
	mapping["printf03"] = stdout.Printf03
	mapping["scanner01"] = stdin.Scanner01
	mapping["map01"] = map_.Map01
	mapping["scope01"] = scope.Scope01
	mapping["async01"] = async.Async01
	mapping["reflection01"] = reflection.Reflection01

	return mapping
}
