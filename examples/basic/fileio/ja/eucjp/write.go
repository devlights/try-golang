package eucjp

import (
	"bytes"
	"io"

	"github.com/devlights/gomy/output"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Write は、EUC-JP なテキストを書き込むサンプルです.
//
// golang.org/x/text/encoding/japanese を利用します.
func Write() error {
	// -------------------------------------------------------------
	// Go で、日本語特有のエンコーディング (sjis や euc-jp など) を扱う場合
	// 以下のパッケージを利用する.
	//
	//   golang.org/x/text/encoding/japanese
	//
	// インストールは、普通に go get するだけ。
	//   go get golang.org/x/text/encoding/japanese
	// go get すると、 golang.org/x/text が依存関係として追加される.
	//
	// japanese.ShiftJIS が sjis, japanese.EUCJP が euc-jp に
	// 対応している。どちらも Encoder と Decoder を持っているので
	// それを transform.NewXXX に渡すことにより、*io.Reader or *io.Writer を
	// 取得することができる。
	//
	// transform.NewReader に渡すのが Decoder
	// transform.NewWriter に渡すのが Encoder
	//
	// となる。
	// -------------------------------------------------------------

	// 書き出し対象データを用意
	var (
		original = []byte("こんにちわ世界")
		buf      = bytes.NewBuffer(original)
	)

	// EUC-JP の エンコーダー 経由で書き出し
	var (
		eucBuf    = new(bytes.Buffer)
		err       error
		encoder   *encoding.Encoder = japanese.EUCJP.NewEncoder()
		eucWriter *transform.Writer = transform.NewWriter(eucBuf, encoder)
	)

	if _, err = io.Copy(eucWriter, buf); err != nil {
		return err
	}

	// 結果確認
	output.Stdoutl("[original]", string(original))
	output.Stdoutl("[euc-jp  ]", eucBuf.String())

	return nil
}
