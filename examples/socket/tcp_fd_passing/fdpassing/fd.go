package fdpassing

import (
	"fmt"
	"net"

	"golang.org/x/sys/unix"
)

// Fd はUnixドメインソケットを使用してファイルディスクリプタをパッシングするための構造体です。
// Unixドメインソケット接続をラップし、ファイルディスクリプタの送受信機能を提供します。
type Fd struct {
	conn *net.UnixConn
}

// NewFd は与えられたUnixドメインソケット接続から新しいFdインスタンスを作成します。
// このコネクションを通じてファイルディスクリプタの送受信が可能になります。
//
// パラメータ:
//   - conn: ファイルディスクリプタの送受信に使用するUnixドメインソケット接続
//
// 戻り値:
//   - 初期化されたFdインスタンスへのポインタ
func NewFd(conn *net.UnixConn) *Fd {
	fd := new(Fd)
	fd.conn = conn
	return fd
}

// Send はファイルディスクリプタをUnixドメインソケット経由で送信します。
// SCM_RIGHTS機能を使用して、プロセス間でファイルディスクリプタを転送します。
//
// パラメータ:
//   - fd: 送信するファイルディスクリプタ
//
// 戻り値:
//   - エラーが発生した場合はエラー、成功した場合はnil
func (me *Fd) Send(fd int) error {
	var (
		dummy  = make([]byte, 1)
		rights = unix.UnixRights(fd)
		err    error
	)
	_, _, err = me.conn.WriteMsgUnix(dummy, rights, nil)
	if err != nil {
		return err
	}

	return nil
}

// Recv はUnixドメインソケット経由でファイルディスクリプタを受信します。
// 送信側から送られたSCM_RIGHTS制御メッセージを解析し、ファイルディスクリプタを取得します。
//
// 戻り値:
//   - 受信したファイルディスクリプタ
//   - エラーが発生した場合はエラー、成功した場合はnil
//     エラーの場合、ファイルディスクリプタは-1が返されます
func (me *Fd) Recv() (int, error) {
	var (
		dummy = make([]byte, 1)
		oob   = make([]byte, unix.CmsgSpace(4))
		flags int
		err   error
	)
	_, _, flags, _, err = me.conn.ReadMsgUnix(dummy, oob)
	if err != nil {
		return -1, err
	}

	if flags&unix.MSG_TRUNC != 0 {
		return -1, fmt.Errorf("control message is truncated")
	}

	var (
		msgs []unix.SocketControlMessage
	)
	msgs, err = unix.ParseSocketControlMessage(oob)
	if err != nil {
		return -1, err
	}

	if len(msgs) != 1 {
		return -1, fmt.Errorf("want: 1 control message; got: %d", len(msgs))
	}

	var (
		fds []int
	)
	fds, err = unix.ParseUnixRights(&msgs[0])
	if err != nil {
		return -1, err
	}

	if len(fds) != 1 {
		return -1, fmt.Errorf("want: 1 fd; got: %d", len(fds))
	}

	return fds[0], nil
}
