package testreader

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

type (
	// UppercaseReader は、読み込んだ文字を大文字に変換するカスタム io.Reader です。
	// ASCIIのみ対応しており、それ以外の文字コードの場合はそのままとなります。
	UppercaseReader struct {
		r io.Reader
		l func(format string, args ...any)
	}
)

func (me *UppercaseReader) Read(p []byte) (n int, err error) {
	if me == nil {
		return
	}

	n, err = me.r.Read(p)
	me.l("len(p)==%2d\tn==%2d\terr==%v", len(p), n, err)

	for i := 0; i < n; i++ {
		if p[i] >= 'a' && p[i] <= 'z' {
			p[i] -= 'a' - 'A'
		}
	}

	return
}

func TestUppercaseReader(t *testing.T) {
	var (
		input    = "hello"
		expected = "HELLO"
		reader   = UppercaseReader{r: strings.NewReader(input), l: t.Logf}
		err      error
	)

	//
	// iotest.TestReader() は、指定された io.Reader の実装が正しく動作するかを
	// テストしてくれる。カスタムな io.Reader を作成した際の動作テストに利用出来る。
	// （バッファである p の値を nil で読んだときの挙動など)
	//
	// 第二引数に指定するのは、期待値となる。
	//
	// - https://pkg.go.dev/testing/iotest@go1.22.6#TestReader
	//
	err = iotest.TestReader(&reader, []byte(expected))
	if err != nil {
		t.Fatal(err)
	}

	/*
	   $ task
	   task: [default] go test -v .
	   === RUN   TestUppercaseReader
	       testreader_test.go:23: len(p)== 0   n== 0   err==<nil>
	       testreader_test.go:23: len(p)== 1   n== 1   err==<nil>
	       testreader_test.go:23: len(p)== 2   n== 2   err==<nil>
	       testreader_test.go:23: len(p)== 3   n== 2   err==<nil>
	       testreader_test.go:23: len(p)== 1   n== 0   err==EOF
	       testreader_test.go:23: len(p)==10   n== 0   err==EOF
	   --- PASS: TestUppercaseReader (0.00s)
	   PASS
	   ok      github.com/devlights/try-golang/examples/singleapp/iotest/testreader    0.002s
	*/

}
