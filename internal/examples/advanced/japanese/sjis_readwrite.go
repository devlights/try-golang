package japanese

import (
	"io/ioutil"
	"os"

	"github.com/devlights/gomy/output"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// SjisReadWrite は、Shift-JISのデータを読み書きするサンプルです.
func SjisReadWrite() error {
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
	dir, err := ioutil.TempDir("", "try-golang")
	if err != nil {
		return err
	}

	defer func() {
		_ = os.RemoveAll(dir)
	}()

	// -------------------------------------------------------------
	// shift-jis データを書き出し
	// -------------------------------------------------------------
	fpath, err := func() (string, error) {
		fp, ioErr := ioutil.TempFile(dir, "sjis")
		if ioErr != nil {
			return "", ioErr
		}

		defer func() {
			_ = fp.Close()
		}()

		writer := transform.NewWriter(fp, japanese.ShiftJIS.NewEncoder())
		_, ioErr = writer.Write([]byte("こんにちわWorld"))
		if ioErr != nil {
			return "", ioErr
		}

		return fp.Name(), nil
	}()

	if err != nil {
		return err
	}

	// -------------------------------------------------------------
	// shift-jis データを読み出し
	// -------------------------------------------------------------
	err = func() error {
		fp, ioErr := os.Open(fpath)
		if ioErr != nil {
			return ioErr
		}

		defer func() {
			_ = fp.Close()
		}()

		reader := transform.NewReader(fp, japanese.ShiftJIS.NewDecoder())
		allData, ioErr := ioutil.ReadAll(reader)
		if ioErr != nil {
			return ioErr
		}

		output.Stdoutl("[sjis]", string(allData))

		return nil
	}()

	return nil
}