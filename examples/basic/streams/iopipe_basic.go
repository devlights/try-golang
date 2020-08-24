package streams

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type (
	_filterWriter struct {
		w io.Writer
	}

	_upperWriter struct {
		w io.Writer
	}
)

var _ io.Writer = (*_filterWriter)(nil)
var _ io.Writer = (*_upperWriter)(nil)

var (
	filterTarget = []byte("wars")
)

func (me *_filterWriter) Write(p []byte) (n int, err error) {
	if bytes.HasPrefix(p, filterTarget) {
		return me.w.Write(p)
	}

	return ioutil.Discard.Write(p)
}

func (me *_upperWriter) Write(p []byte) (n int, err error) {
	return me.w.Write(bytes.ToUpper(p))
}

// IoPipeBasic -- io.Pipe の基本的なサンプルです.
//
// REFERENCES:
//   - https://www.geeksforgeeks.org/io-pipe-function-in-golang-with-examples/
//   - https://medium.com/eureka-engineering/file-uploads-in-go-with-io-pipe-75519dfa647b
func IoPipeBasic() error {
	// ----------------------------------------------------------------
	// io.Pipe() は、インメモリの同期パイプを生成してくれる。
	// io.Pipe() は、片方が io.Reader を必要としていて
	// もう片方が、io.Writer を必要としてる時に、その2つを接続することが出来る。
	//
	// 通常、このような場合は中間バッファを作成し結果を蓄積し
	// 出力先の io.Writer に出力することが多い。
	// しかし、大容量のデータを中間バッファに蓄積する場合は
	// それがメモリを大量に消費する原因となる可能性がある。
	//
	// io.Pipe() で取得できる *PipeReader と *PipeWriter は
	// 同期しており、かつ、バッファリング無しとなっているので
	// *PipeWriter に書き込んだデータは、即座に *PipeReader 側に渡る。
	//
	// 注意点として、同期しているストリームなので
	// 片方だけ利用するということは出来ない。デッドロックが発生する。
	// *PipeReaderで読むためには、必ず *PipeWriter 側で書き込みを行っていないといけない
	// 逆も然り。
	//
	// なので、io.Pipe() を利用する場合は、ゴルーチンが必要となる。
	// ----------------------------------------------------------------
	var (
		pr *io.PipeReader
		pw *io.PipeWriter
	)

	pr, pw = io.Pipe()

	// 書き込み側を非同期で実行
	// io.Pipe() から取得できる2つのストリームは同期しているので
	// 同時に扱わないとデッドロックする.
	go func() {
		defer pw.Close()

		f, err := os.Open("/usr/share/dict/words")
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			fmt.Fprintf(pw, "%s\n", scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()

	// 読み込み
	// *PipeReader側で読み込むと*PipeWriter側のデータはクリアされる
	upperW := &_upperWriter{os.Stdout}
	filterW := &_filterWriter{upperW}
	io.Copy(filterW, pr)

	return nil
}
