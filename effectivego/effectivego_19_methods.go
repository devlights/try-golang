package effectivego

import (
	"fmt"

	"github.com/devlights/try-golang/lib/output"
)

type (
	valueReceiverIntSlice   []int
	pointerReceiverIntSlice []int
	limitedByteSlice        []byte
)

func (s valueReceiverIntSlice) Append(data ...int) []int {
	return append(s, data...)
}

func (s *pointerReceiverIntSlice) Append(data ...int) {
	v := *s
	v = append(v, data...)
	*s = v
}

func (s *limitedByteSlice) Write(data []byte) (n int, err error) {
	v := *s
	v = append(v, data...)
	v = v[:5]
	*s = v

	return len(v), nil
}

// Effective Go - Methods の 内容についてのサンプルです。
func Methods() error {
	/*
		https://golang.org/doc/effective_go.html#methods

		- Value Receiver は、レシーバに値がコピーされた渡される
		  - なので、レシーバを変更しても元の値には影響がない
		- Pointer Receiver は、レシーバに値のポインタがコピーされて渡される
		  - なので、レシーバを変更すると元の値にも影響する
	*/
	s1 := valueReceiverIntSlice{1, 2, 3, 4, 5}
	s2 := pointerReceiverIntSlice{1, 2, 3, 4, 5}

	s3 := s1.Append(99, 100)
	output.Stdoutl("(1)", s1, s3)

	// ポインタレシーバの呼び出しは、本来 (&s2).Append() としないといけないが
	// Goでは利便性を重視して、コンパイラが呼び出しを置き換えてくれるようになっている
	// なので、以下のように s2.Append() としてもポインタレシーバのメソッドが呼び出される
	s2.Append(99, 100)
	output.Stdoutl("(2)", s2)

	(&s2).Append(101, 102)
	output.Stdoutl("(3)", s2)

	// io.Writerの実装
	//  ポインタレシーバのメソッドとしてWrite()を定義しているので
	//  渡す際にポインタで渡す必要がある
	var b limitedByteSlice
	n, err := fmt.Fprint(&b, "helloworld")
	output.Stdoutl("(4)", string(b), n, err)

	return nil
}
