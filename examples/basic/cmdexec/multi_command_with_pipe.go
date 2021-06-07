package cmdexec

import (
	"bytes"
	"io"
	"os/exec"

	"github.com/devlights/gomy/output"
)

// MultiCommandWithPipe は、複数の (*exec.Cmd) をパイプストリームで繋いで実行するサンプルです.
func MultiCommandWithPipe() error {
	var (
		// $ find . -name "*.go" | grep -v -E ".*(doc|examples)\.go" | wc -l
		cmds = []*exec.Cmd{
			exec.Command("find", ".", "-name", "*.go"),
			exec.Command("grep", "-v", "-E", ".*(doc|examples)\\.go"),
			exec.Command("wc", "-l"),
		}
	)

	var (
		in  = new(bytes.Buffer) // 先頭ステージの入力
		out = new(bytes.Buffer) // 末尾ステージの出力
	)

	// 先頭と末尾の入出力を設定
	cmds[0].Stdin = in
	cmds[len(cmds)-1].Stdout = out

	// 各ステージの出力を次のコマンドの入力につなぐ
	for i := 0; i < len(cmds)-1; i++ {
		var (
			curr    = cmds[i]
			next    = cmds[i+1]
			currOut io.ReadCloser
			err     error
		)

		if currOut, err = curr.StdoutPipe(); err != nil {
			return err
		}

		next.Stdin = currOut
	}

	// 実行
	for _, c := range cmds {
		if err := c.Start(); err != nil {
			return err
		}
	}

	// 終了待機
	for _, c := range cmds {
		if err := c.Wait(); err != nil {
			return err
		}
	}

	// 結果出力
	output.Stdoutl("[result]", out.String())

	return nil
}
