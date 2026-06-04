// 出力先を os.Stdout から *bufio.Writer に変更したサンプルです。
//
// REFERENCES
//   - https://stackoverflow.com/questions/64638136/performance-issues-while-reading-a-file-line-by-line-with-bufio-newscanner
//   - https://yourbasic.org/golang/temporary-file-directory/
//   - https://pkg.go.dev/os

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devlights/gomy/output"
)

func main() {
	var (
		fpath = "/tmp/try-golang-big.txt"
		file  *os.File
		err   error
	)

	if file, err = os.Open(fpath); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	var (
		scanner = bufio.NewScanner(file)
		writer  = bufio.NewWriter(os.Stdout)
		start   = time.Now()
	)

	for scanner.Scan() {
		fmt.Fprintln(writer, scanner.Text())
	}

	output.Stderrl("[Elapsed]", time.Since(start))
}
