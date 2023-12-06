package sjis

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/devlights/gomy/output"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Read は、Shift-JISなテキストを読み込むサンプルです.
//
// golang.org/x/text/encoding/japanese を利用します.
func Read() error {
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
	const (
		shell = "/usr/bin/bash"
		fpath = "/tmp/sjis.txt"
	)

	// 先にサンプルとなる Shift-JIS エンコーディング のファイルを作成しておく
	var (
		cmd    *exec.Cmd
		err    error
		cmdTxt = fmt.Sprintf("echo 'こんにちわ世界' | nkf -sd > %s", fpath)
	)

	cmd = exec.Command(shell, "-c", cmdTxt)
	if err = cmd.Run(); err != nil {
		return err
	}

	// 対象ファイルのデータを読み出し
	var (
		buf       *bytes.Buffer
		origBytes []byte
	)

	if origBytes, err = os.ReadFile(fpath); err != nil {
		return err
	}
	buf = bytes.NewBuffer(origBytes)

	// Shift-JIS で デコード
	var (
		decoder    *encoding.Decoder = japanese.ShiftJIS.NewDecoder()
		sjisReader *transform.Reader = transform.NewReader(buf, decoder)
		sjisBytes  []byte
	)

	if sjisBytes, err = io.ReadAll(sjisReader); err != nil {
		return err
	}

	// 結果出力
	output.Stdoutl("[original]", string(origBytes))
	output.Stdoutl("[shiftjis]", string(sjisBytes))

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: fileio_japanese_sjis_read

	   [Name] "fileio_japanese_sjis_read"
	   [original]           ����ɂ��퐢�E

	   [shiftjis]           こんにちわ世界



	   [Elapsed] 3.89531ms
	*/

}
