package fileio

import (
	"bufio"
	"os"

	"github.com/devlights/gomy/output"
)

// StdinWithScanner -- os.Stdin と bufio.Scanner のサンプルです.
func StdinWithScanner() error {
	// --------------------------------------------------
	// go で 標準入力 を 扱う際に最もやりやすいのは
	// bufio.Scanner を利用する方法。
	//
	// bufio.NewScanner() の引数に os.Stdin を渡すことで
	// 標準入力のハンドリングができる。
	// --------------------------------------------------
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "quit" {
			break
		}

		output.Stdoutl("[text]", txt)
	}

	return scanner.Err()
}
