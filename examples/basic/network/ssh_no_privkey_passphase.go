package network

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/devlights/gomy/output"
	"golang.org/x/crypto/ssh"
)

type (
	// returnCode -- 処理結果
	returnCode int
)

const (
	success                   returnCode = iota // 成功
	homedirNotFound                             // $HOME が展開出来なかった
	readErrSSHPrivateKey                        // 秘密鍵読み取り中にエラー
	parseErrSSHPrivateKey                       // 秘密鍵の解析中にエラー
	parseErrPublicKey                           // 公開鍵の解析中にエラー
	connErrSSHClient                            // SSH接続中にエラー
	canNotCreateNewSSHSession                   // SSHにてセッションを生成中にエラー
	execErrInSSHSession                         // SSHにてコマンドを実行中にエラー
)

const (
	command = "cat /etc/os-release | head -n 2"
)

// 環境変数より取得する情報
var (
	sshUser string // SSH ユーザ名
	sshPass string // SSH パスワード
	sshHost string // SSH リモートホスト (xxx.xxx.xxx.xxx:port)
	hostKey string // リモートホストの公開鍵 (ssh-keyscan の 結果)(e.g: xxx.xxx.xxx.xxx ecdsa-sha2-nistp256 xxxxxxxxxxx)
)

// SSHNoPrivKeyPassPhase -- 秘密鍵のパスフレーズ無しのSSH接続サンプルです.
//
// 本サンプルを実行するには、以下の環境変数が設定されている必要があります。
//
// - $SSH_USER: SSHユーザ名
//
// - $SSH_PASS: SSHパスワード
//
// - $SSH_HOST: 接続先を xxx.xxx.xxx.xxx:port の形式で
//
// - $HOST_KEY: リモートサーバの公開鍵 (ssh-keyscan -4 -t ecdsa remote-host の結果)
//
//noinspection GoErrorStringFormat
func SSHNoPrivKeyPassPhase() error {
	sshUser = os.ExpandEnv("$SSH_USER")
	sshPass = os.ExpandEnv("$SSH_PASS")
	sshHost = os.ExpandEnv("$SSH_HOST")
	hostKey = os.ExpandEnv("$HOST_KEY")

	if sshUser == "" {
		return errors.New("$SSH_USER が 設定されていません.")
	}

	if sshPass == "" {
		return errors.New("$SSH_PASS が 設定されていません.")
	}

	if sshHost == "" {
		return errors.New("$SSH_HOST が 設定されていません.")
	}

	if hostKey == "" {
		return errors.New("$HOST_KEY が 設定されていません.")
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
	// なお、今回は ssh-keygen で パスワード無しの秘密鍵を作ったとする。
	// 秘密鍵にパスワードを付与している場合の処理は別のサンプルにて。
	//
	// 以下、それぞれのケースで同じことをしている。
	// -------------------------------------------------------------

	// -------------------------------------------------------------
	// (1) SSH 接続してコマンド実行
	//   - パスワード認証
	//   - リモートサーバの公開鍵を検証しない
	// -------------------------------------------------------------
	fmt.Println("RUN (1)")
	if ret := sshWithPasswordWithInsecureHostKey(); ret != success {
		return fmt.Errorf("[error] RUN (1) returnCode:%v", ret)
	}

	// -------------------------------------------------------------
	// (2) SSH 接続してコマンド実行
	//   - パスワード認証
	//   - リモートサーバの公開鍵を検証する
	// -------------------------------------------------------------
	fmt.Println("RUN (2)")
	if ret := sshWithPasswordWithFixedHostKey(); ret != success {
		return fmt.Errorf("[error] RUN (2) returnCode:%v", ret)
	}

	// -------------------------------------------------------------
	// (3) SSH 接続してコマンド実行
	//   - 鍵認証
	//   - リモートサーバの公開鍵を検証しない
	// -------------------------------------------------------------
	fmt.Println("RUN (3)")
	if ret := sshWithKeyFileWithInsecureHostKey(); ret != success {
		return fmt.Errorf("[error] RUN (3) returnCode:%v", ret)
	}

	// -------------------------------------------------------------
	// (4) SSH 接続してコマンド実行
	//   - 鍵認証
	//   - リモートサーバの公開鍵を検証する
	// -------------------------------------------------------------
	fmt.Println("RUN (4)")
	if ret := sshWithKeyFileWithFixedHostKey(); ret != success {
		return fmt.Errorf("[error] RUN (4) returnCode:%v", ret)
	}

	return nil
}

// (1) SSH 接続してコマンド実行
//   - パスワード認証
//   - リモートサーバの公開鍵を検証しない
func sshWithPasswordWithInsecureHostKey() returnCode {
	// -------------------------------------------
	// SSH の 接続設定 を構築
	//
	config := &ssh.ClientConfig{
		// SSH ユーザ名
		User: sshUser,
		// 認証方式
		Auth: []ssh.AuthMethod{
			// パスワード認証
			ssh.Password(sshPass),
		},
		// リモートサーバの公開鍵を検証しない
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// -------------------------------------------
	// SSH で 接続
	//
	conn, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		output.Stderrl("[ssh.Dial]", err)
		return connErrSSHClient
	}

	// -------------------------------------------
	// セッションを開いて、コマンドを実行
	//
	sess, err := conn.NewSession()
	if err != nil {
		output.Stderrl("[conn.NewSession]", err)
		return canNotCreateNewSSHSession
	}

	defer func() {
		e := sess.Close()
		if e != nil {
			output.Stderrl("[sess.Close]", e)
		}
	}()

	// リモートサーバでのコマンド実行結果をローカルの標準出力と標準エラーへ流す
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	if err = sess.Run(command); err != nil {
		output.Stderrl("[sess.Run]", err)
		return execErrInSSHSession
	}

	return success
}

// (2) SSH 接続してコマンド実行
//   - パスワード認証
//   - リモートサーバの公開鍵を検証する
func sshWithPasswordWithFixedHostKey() returnCode {
	// -------------------------------------------
	// リモートサーバ の 公開鍵 を得る
	//
	_, _, pubKey, _, _, err := ssh.ParseKnownHosts([]byte(hostKey))
	if err != nil {
		output.Stderrl("[ssh.ParseKnownHosts]", err)
		return parseErrPublicKey
	}

	// -------------------------------------------
	// SSH の 接続設定 を構築
	//
	config := &ssh.ClientConfig{
		// SSH ユーザ名
		User: sshUser,
		// 認証方式
		Auth: []ssh.AuthMethod{
			// パスワード認証
			ssh.Password(sshPass),
		},
		// リモートサーバの公開鍵を検証
		HostKeyCallback: ssh.FixedHostKey(pubKey),
	}

	// -------------------------------------------
	// SSH で 接続
	//
	conn, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		output.Stderrl("[ssh.Dial]", err)
		return connErrSSHClient
	}

	// -------------------------------------------
	// セッションを開いて、コマンドを実行
	//
	sess, err := conn.NewSession()
	if err != nil {
		output.Stderrl("[conn.NewSession]", err)
		return canNotCreateNewSSHSession
	}

	defer func() {
		e := sess.Close()
		if e != nil {
			output.Stderrl("[sess.Close]", e)
		}
	}()

	// リモートサーバでのコマンド実行結果をローカルの標準出力と標準エラーへ流す
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	if err = sess.Run(command); err != nil {
		output.Stderrl("[sess.Run]", err)
		return execErrInSSHSession
	}

	return success
}

// (3) SSH 接続してコマンド実行
//   - 鍵認証
//   - リモートサーバの公開鍵を検証しない
func sshWithKeyFileWithInsecureHostKey() returnCode {
	// -------------------------------------------
	// $HOME/.ssh/id_rsa からデータ読み取り
	//
	homeDir, err := os.UserHomeDir()
	if err != nil {
		output.Stderrl("[os.UserHomeDir]", err)
		return homedirNotFound
	}

	sshPrivKeyFile := filepath.Join(homeDir, ".ssh/id_rsa")
	privKey, err := ioutil.ReadFile(sshPrivKeyFile)
	if err != nil {
		output.Stderrl("[ioutil.ReadFile]", err)
		return readErrSSHPrivateKey
	}

	// -------------------------------------------
	// 秘密鍵を渡して Signer を取得
	//
	signer, err := ssh.ParsePrivateKey(privKey)
	if err != nil {
		output.Stderrl("[ssh.ParsePrivateKey]", err)
		return parseErrSSHPrivateKey
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
		// リモートサーバの公開鍵を検証しない
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// -------------------------------------------
	// SSH で 接続
	//
	conn, err := ssh.Dial("tcp", sshHost, config)
	if err != nil {
		output.Stderrl("[ssh.Dial]", err)
		return connErrSSHClient
	}

	// -------------------------------------------
	// セッションを開いて、コマンドを実行
	//
	sess, err := conn.NewSession()
	if err != nil {
		output.Stderrl("[conn.NewSession]", err)
		return canNotCreateNewSSHSession
	}

	defer func() {
		e := sess.Close()
		if e != nil {
			output.Stderrl("[sess.Close]", e)
		}
	}()

	// リモートサーバでのコマンド実行結果をローカルの標準出力と標準エラーへ流す
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	if err = sess.Run(command); err != nil {
		output.Stderrl("[sess.Run]", err)
		return execErrInSSHSession
	}

	return success
}

// (4) SSH 接続してコマンド実行
//   - 鍵認証
//   - リモートサーバの公開鍵を検証する
func sshWithKeyFileWithFixedHostKey() returnCode {
	// -------------------------------------------
	// $HOME/.ssh/id_rsa からデータ読み取り
	//
	homeDir, err := os.UserHomeDir()
	if err != nil {
		output.Stderrl("[os.UserHomeDir]", err)
		return homedirNotFound
	}

	sshPrivKeyFile := filepath.Join(homeDir, ".ssh/id_rsa")
	privKey, err := ioutil.ReadFile(sshPrivKeyFile)
	if err != nil {
		output.Stderrl("[ioutil.ReadFile]", err)
		return readErrSSHPrivateKey
	}

	// -------------------------------------------
	// 秘密鍵を渡して Signer を取得
	//
	signer, err := ssh.ParsePrivateKey(privKey)
	if err != nil {
		output.Stderrl("[ssh.ParsePrivateKey]", err)
		return parseErrSSHPrivateKey
	}

	// -------------------------------------------
	// リモートサーバ の 公開鍵 を得る
	//
	_, _, pubKey, _, _, err := ssh.ParseKnownHosts([]byte(hostKey))
	if err != nil {
		output.Stderrl("[ssh.ParseKnownHosts]", err)
		return parseErrPublicKey
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
		output.Stderrl("[ssh.Dial]", err)
		return connErrSSHClient
	}

	// -------------------------------------------
	// セッションを開いて、コマンドを実行
	//
	sess, err := conn.NewSession()
	if err != nil {
		output.Stderrl("[conn.NewSession]", err)
		return canNotCreateNewSSHSession
	}

	defer func() {
		e := sess.Close()
		if e != nil {
			output.Stderrl("[sess.Close]", e)
		}
	}()

	// リモートサーバでのコマンド実行結果をローカルの標準出力と標準エラーへ流す
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	if err = sess.Run(command); err != nil {
		output.Stderrl("[sess.Run]", err)
		return execErrInSSHSession
	}

	return success
}
