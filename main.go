package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/devlights/try-golang/advanced/async"
	"github.com/devlights/try-golang/advanced/closure"
	"github.com/devlights/try-golang/advanced/reflection"
	"github.com/devlights/try-golang/basic/array_"
	"github.com/devlights/try-golang/basic/comments"
	"github.com/devlights/try-golang/basic/helloworld"
	"github.com/devlights/try-golang/basic/import_"
	"github.com/devlights/try-golang/basic/interface_"
	"github.com/devlights/try-golang/basic/io_"
	"github.com/devlights/try-golang/basic/map_"
	"github.com/devlights/try-golang/basic/scope"
	"github.com/devlights/try-golang/basic/slice_"
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"github.com/devlights/try-golang/basic/struct_"
	"log"
	"os"
	"strings"
)

func main() {
	var (
		onetime = flag.Bool("onetime", false, "run only one time")
	)

	flag.Parse()

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

		if *onetime {
			break
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
	mapping["println01"] = stdout.Println01
	mapping["scanner01"] = stdin.Scanner01
	mapping["map01"] = map_.Map01
	mapping["scope01"] = scope.Scope01
	mapping["async01"] = async.Async01
	mapping["reflection01"] = reflection.Reflection01
	mapping["import01"] = import_.Import01
	mapping["fileio01"] = io_.FileIo01
	mapping["fileio02"] = io_.FileIo02
	mapping["fileio03"] = io_.FileIo03
	mapping["fileio04"] = io_.FileIo04
	mapping["interface01"] = interface_.Interface01
	mapping["struct01"] = struct_.Struct01
	mapping["array01"] = array_.Array01
	mapping["slice01"] = slice_.Slice01
	mapping["slice02"] = slice_.Slice02
	mapping["slice03"] = slice_.Slice03
	mapping["slice04"] = slice_.Slice04
	mapping["slice05"] = slice_.Slice05
	mapping["comment01"] = comments.Comment01
	mapping["closure01"] = closure.Closure01

	return mapping
}
