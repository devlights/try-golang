package errreader

import (
	"errors"
	"strings"
	"syscall"
	"testing"
	"testing/iotest"
)

func TestReadOk(t *testing.T) {
	var (
		s   = "hello world"
		r   = strings.NewReader(s)
		p   = make([]byte, len(s))
		err error
	)

	err = read(r, p)
	if err != nil {
		t.Fatal(err)
	}

	if string(p) != s {
		t.Fatalf("[want] equal\t[got] not equal (%s, %s)", string(p), s)
	}
}

func TestReadTooManyEAGAIN(t *testing.T) {
	//
	// iotest.ErrReader() は、指定したエラーを返す io.Reader を返してくれる。
	// io.Readerを使って何らかの処理を行う関数などを実装している場合のエラーテストに便利。
	// (例えば、ノンブロッキング処理をしている実装で、ずっとEAGAINが返ってくる場合のテストなど)
	//
	// - https://pkg.go.dev/testing/iotest@go1.23.0#ErrReader
	//
	var (
		s = "hello world"
		r = iotest.ErrReader(syscall.EAGAIN)
		p = make([]byte, len(s))
	)

	err := read(r, p)
	if err == nil {
		t.Fatal("[want] err \t[got] nil")
	}

	if !errors.Is(err, ErrTooManyEAGAIN) {
		t.Fatalf("[want] ErrRetryOver\t[got] %v", err)
	}
}
