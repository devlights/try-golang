package network

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/devlights/gomy/output"
	"golang.org/x/crypto/ssh"
)

// SSHWithPrivKeyPassPhrase -- 秘密鍵のパスフレーズありのSSH接続サンプルです.
//
// 本サンプルを実行するには、以下の環境変数が設定されている必要があります。
//
// - $SSH_USER: SSHユーザ名
//
// - $SSH_PASS: 秘密鍵のパスワード
//
// - $SSH_HOST: 接続先を xxx.xxx.xxx.xxx:port の形式で
//
// - $HOST_KEY: リモートサーバの公開鍵 (ssh-keyscan -4 -t ecdsa remote-host の結果)
//
// REFERENCES:
//   - https://stackoverflow.com/questions/60879023/getting-eof-as-error-in-golang-ssh-session-close
//
//noinspection GoErrorStringFormat
func SSHWithPrivKeyPassPhrase() error {
	sshUser = os.ExpandEnv("$SSH_USER")
	sshPass = os.ExpandEnv("$SSH_PASS")
	sshHost = os.ExpandEnv("$SSH_HOST")
	hostKey = os.ExpandEnv("$HOST_KEY")

	if sshUser == "" {
		return errors.New("$SSH_USER が 設定されていません")
	}

	if sshPass == "" {
		return errors.New("$SSH_PASS が 設定されていません")
	}

	if sshHost == "" {
		return errors.New("$SSH_HOST が 設定されていません")
	}

	if hostKey == "" {
		return errors.New("$HOST_KEY が 設定されていません")
	}

	// -------------------------------------------------------------
	// GO で ssh を扱う場合 golang.org/x/crypto/ssh を使う
	//
	// 標準パッケージには入っていないので利用する場合は go get する.
	//   $ go get -v -u golang.org/x/crypto/ssh
	//
	// SSH で接続する場合、大きく分けて
	//   1. パスワード認証
	//   2. 鍵認証
	// の２つがある。
	//
	// また、リモートサーバの 公開鍵 を
	//   1. 検証する
	//   2. 検証しない
	// の２つがある。
	//
	// なお、今回は ssh-keygen で パスワード付きの秘密鍵を作ったとする。
	// -------------------------------------------------------------

	// -------------------------------------------------------------
	// SSH 接続してコマンド実行
	//   - 鍵認証
	//   - リモートサーバの公開鍵を検証する
	// -------------------------------------------------------------

	// -------------------------------------------
	// $HOME/.ssh/id_rsa からデータ読み取り
	//
	homeDir, err := os.UserHomeDir()
	if err != nil {
		output.Stderrl("[error]", err)
		return err
	}

	sshPrivKeyFile := filepath.Join(homeDir, ".ssh/id_rsa")
	privKey, err := os.ReadFile(sshPrivKeyFile)
	if err != nil {
		output.Stderrl("[error]", err)
		return err
	}

	// -------------------------------------------
	// 秘密鍵を渡して Signer を取得
	//
	signer, err := ssh.ParsePrivateKeyWithPassphrase(privKey, []byte(sshPass))
	if err != nil {
		output.Stderrl("[error]", err)
		return err
	}

	// -------------------------------------------
	// リモートサーバ の 公開鍵 を得る
	//
	_, _, pubKey, _, _, err := ssh.ParseKnownHosts([]byte(hostKey))
	if err != nil {
		output.Stderrl("[error]", err)
		return err
	}

	// -------------------------------------------
	// SSH の 接続設定 を構築
	//
	config := &ssh.ClientConfig{
		// SSH ユーザ名
		User: sshUser,
		// 認証方式
		Auth: []ssh.AuthMethod{
			// 鍵認証
			ssh.PublicKeys(signer),
		},
		// リモートサーバの公開鍵を検証
		HostKeyCallback: ssh.FixedHostKey(pubKey),
	}

	// -------------------------------------------
	// SSH で 接続
	//
	conn, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		output.Stderrl("[error]", err)
		return err
	}

	// -------------------------------------------
	// セッションを開いて、コマンドを実行
	//
	sess, err := conn.NewSession()
	if err != nil {
		output.Stderrl("[error]", err)
		return err
	}
	defer sess.Close()

	// リモートサーバでのコマンド実行結果をローカルの標準出力と標準エラーへ流す
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	if err = sess.Run(command); err != nil {
		output.Stderrl("[error]", err)
		return err
	}

	return nil
}
