package scannerop

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Basic -- bufio.Scannerの基本的な使い方についてのサンプルです。
func Basic() error {
	fmt.Print("Enter: ")

	// 標準入力から入力を読み込むには以下のようにする
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		if strings.ToLower(text) == "quit" {
			break
		}

		fmt.Println(text)
		fmt.Print("Enter: ")
	}

	fmt.Println("END")

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: scannerop_basic

	   [Name] "scannerop_basic"
	   Enter: helloworld
	   helloworld
	   Enter: quit
	   END


	   [Elapsed] 4.979184142s
	*/

}
