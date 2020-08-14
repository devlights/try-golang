package iointerfaces

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/devlights/gomy/output"
)

type (
	// _readcloserimpl -- サンプル用 io.ReadCloser の実装
	_readcloserimpl struct {
		reader io.Reader
	}
)

// Impl check
var _ io.ReadCloser = (*_readcloserimpl)(nil)

// Impl -- io.Reader#Read
func (r *_readcloserimpl) Read(p []byte) (int, error) {
	return r.reader.Read(p)
}

// Impl -- io.Closer#Close
func (r *_readcloserimpl) Close() error {
	output.Stdoutl("[_readcloserimpl]", "Close() called")
	return nil
}

// UsingNopCloser -- ioutil.NopCloser についてのサンプルです.
//
// REFERENCES:
//   - https://golang.org/pkg/io/ioutil/#NopCloser
func UsingNopCloser() error {
	// ------------------------------------------------------------
	// ioutil.NopCloser は、Closeが呼ばれても何もしない
	// io.ReadCloser を返す.
	//
	// 処理内部で Close メソッドが呼ばれてしまうが
	// そのタイミングで Close はしてほしくない場合などで利用できる.
	// ------------------------------------------------------------
	var (
		b = []byte("hello")
		r = bytes.NewReader(b)
	)

	reader := &_readcloserimpl{reader: r}
	nopcloser := ioutil.NopCloser(reader)

	buf := make([]byte, 3)
	_, _ = nopcloser.Read(buf)
	output.Stdoutl("[nopcloser.Read]", buf)

	output.Stdoutl(">>", "before NopCloser.Close")
	_ = nopcloser.Close()
	output.Stdoutl(">>", "after  NopCloser.Close")

	_ = reader.Close()

	return nil
}
