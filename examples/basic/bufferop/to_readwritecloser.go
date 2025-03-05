package bufferop

import (
	"bytes"
	"io"
	"os"
)

// ToReadWriteCloser は、bytes.Buffer を io.ReadWriteCloser に変換するサンプルです.
//
// bytes.Bufferは、既に io.Reader と io.Writer を実装していますが
// io.Closer は実装していません。そのため、そのままでは
// os.File や net.Conn のように io.ReadWriteCloser としては利用できません。
//
// bytes.Bufferをラップし、Closeメソッドを空実装した型を別途用意すれば良いのですが面倒です。
// リミテッドな使い方（特定の関数に渡したいだけ）の場合は、匿名構造体と io.NopCloser を利用した方が楽です。
//
// # REFERENCES
//   - https://zenn.dev/nobonobo/articles/297dc5cbc554d6
func ToReadWriteCloser() error {
	// funcs
	var (
		// 変換
		c = func(b *bytes.Buffer) io.ReadWriteCloser {
			// 匿名構造体を使って、io.ReadWriteCloser の実装を用意。
			//   io.ReadWriteCloser は io.Reader, io.Writer, io.Closer の合成インターフェースとなる。
			//   なので、その３つを実装していれば良いことになる。
			return &struct {
				io.ReadCloser // io.Reader と io.Closer
				io.Writer     // io.Writer
			}{
				io.NopCloser(b), // io.ReaderはBufferが実装済み。io.Closerは空実装。
				b,               // io.WriterはBufferが実装済み。
			}
		}
		// 書き込み
		w = func(rwc io.ReadWriteCloser) {
			_, _ = rwc.Write([]byte("hello\n"))
		}
		// 読み込み
		r = func(rwc io.ReadWriteCloser) {
			io.Copy(os.Stdout, rwc)
		}
	)

	var (
		b   bytes.Buffer
		rwc = c(&b)
	)
	w(rwc)
	r(rwc)

	return nil
}
