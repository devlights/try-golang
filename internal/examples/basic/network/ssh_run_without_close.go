package network

import (
	"fmt"
	"io"
	"os"

	"github.com/devlights/gomy/output"
	"golang.org/x/crypto/ssh"
)

// SSHSessionCloseAfterRun -- ssh.Run() を呼んだ後に ssh.Close() を呼ぶと io.EOF が返却されることを確認するサンプルです.
//
// REFERENCES:
//   - https://stackoverflow.com/questions/60879023/getting-eof-as-error-in-golang-ssh-session-close
func SSHSessionCloseAfterRun() error {
	// -----------------------------------------------------------------------------------------
	// ssh.Run() した後に ssh.Close() すると io.EOF が返却される件
	//
	// 理由としては、ssh.Run() を実行すると内部でセッションを Close しているため。
	// 適切に セッション は閉じられているので、このエラーは無視しても良い。
	// -----------------------------------------------------------------------------------------
	var (
		sshUser = os.ExpandEnv("${SSH_USER}")
		sshPass = os.ExpandEnv("${SSH_PASS}")
		sshHost = os.ExpandEnv("${SSH_HOST}")
	)

	if sshUser == "" {
		return fmt.Errorf("[error] %s が設定されていません", "${SSH_USER}")
	}

	if sshPass == "" {
		return fmt.Errorf("[error] %s が設定されていません", "${SSH_PASS}")
	}

	if sshHost == "" {
		return fmt.Errorf("[error] %s が設定されていません", "${SSH_HOST}")
	}

	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		output.Stderrl("[ssh.Dial]", err)
		return err
	}

	sess, err := conn.NewSession()
	if err != nil {
		output.Stderrl("[conn.NewSession]", err)
		return err
	}

	defer func() {
		// see: https://stackoverflow.com/questions/60879023/getting-eof-as-error-in-golang-ssh-session-close
		//
		// sess.Run() は、コマンドを実行後にセッションをcloseしている.
		// そのため、sess.Run() 後に sess.Close() を呼ぶと io.EOF が返却される.
		// 既にセッションは sess.Run() で適切にクローズされているため、このエラーは無視しても良い
		e := sess.Close()
		if e != nil && e != io.EOF {
			output.Stderrl("[sess.Close]", e)
		}
	}()

	// リモートサーバでのコマンド実行結果をローカルの標準出力と標準エラーへ流す
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	if err = sess.Run(command); err != nil {
		output.Stderrl("[sess.Run]", err)
		return err
	}

	return nil
}
