// 標準入力からデータを読み込むサンプル
//
// REFERENCES:: http://bit.ly/2W3DCM4
package stdin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Scanner01() error {
	fmt.Print("Enter example-name: ")

	// 標準入力から入力を読み込むには以下のようにする
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		if strings.ToLower(text) == "quit" {
			break
		}

		fmt.Println(text)
		fmt.Print("Enter example-name: ")
	}

	fmt.Println("END")

	return nil
}
