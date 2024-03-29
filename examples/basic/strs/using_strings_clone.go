package strs

import (
	"os/exec"
	"strings"

	"github.com/devlights/gomy/output"
)

// UsingStringsClone は、Go 1.18 で追加された strings.Clone() のサンプルです。
//
// # REFERENCES
//   - https://pkg.go.dev/strings@go1.21.4#Clone
func UsingStringsClone() error {
	//
	// strings.Clone() は、Go 1.18 で追加された関数。
	// 新たなメモリ割り当ても行われるので、ディープコピーされるイメージ。
	// 部分文字列を特定のストアに対して保持するようなシチュエーションで利用できる。
	// (Goの標準コンパイラでは、現状元の文字列と部分文字列は同じメモリを共有するようになっているため
	//  数が多い、または、文字列のサイズが巨大な場合などにディープコピーしておかないとメモリが残ることなる）
	//

	const (
		NumItems = 100
	)

	var (
		l = make([]string, NumItems)
	)

	// 1024バイトのランダム文字列を1000個用意
	for i := 0; i < NumItems; i++ {
		var (
			c   *exec.Cmd
			o   []byte
			err error
		)

		c = exec.Command("/bin/bash", "-c", "openssl rand -base64 1024 | tr -d '\n'")
		if o, err = c.Output(); err != nil {
			return err
		}

		l[i] = string(o)
	}

	//
	// 各文字列の先頭５文字分を識別子として保持しておく仕様があるとする
	//

	var (
		store1 = make([]string, NumItems)
		store2 = make([]string, NumItems)
	)

	// 部分文字列を取り出し、そのまま保持
	// この場合、元文字列と部分文字列は同じメモリを共有している可能性があるため
	// 場合によっては、５バイト分だけじゃなく、文字列全部がメモリに残ったままとなる
	for i := 0; i < NumItems; i++ {
		store1[i] = l[i][:5]
	}

	// 部分文字列を取り出し、クローンしてから保持
	// Go 1.18 で追加された strings.Clone() を利用することで、新たな割当が行われた状態で
	// 文字列がクローンされる。なので元文字列全部がメモリに残ることはなくなる
	for i := 0; i < NumItems; i++ {
		store2[i] = strings.Clone(l[i][:5])
	}

	for i := 0; i < 3; i++ {
		output.Stdoutf("[store1]", "%d: %s\n", i, store1[i])
		output.Stdoutf("[store2]", "%d: %s\n", i, store2[i])
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: string_using_clone

	   [Name] "string_using_clone"
	   [store1]             0: rq0Py
	   [store2]             0: rq0Py
	   [store1]             1: VtRco
	   [store2]             1: VtRco
	   [store1]             2: MSL1F
	   [store2]             2: MSL1F


	   [Elapsed] 540.617051ms
	*/

}
