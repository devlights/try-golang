//go:build linux

package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sys/unix"
)

// SocketFd は、ソケットファイルディスクリプタを表します。
type SocketFd int

// Readable は、select(2)を呼び出し読み込み可能かどうかを判定します。
func (me SocketFd) Readable(sec, usec time.Duration) (bool, error) {
	fd := int(me)
	if fd < 0 || fd >= unix.FD_SETSIZE {
		return false, fmt.Errorf("invalid file descriptor: out of range %d (FD_SETSIZE = %d)", fd, unix.FD_SETSIZE)
	}

	rfds := &unix.FdSet{}
	rfds.Zero()
	rfds.Set(fd)

	timeout := &unix.Timeval{
		Sec:  int64(sec.Seconds()),
		Usec: usec.Microseconds(),
	}

	n, err := unix.Select(fd+1, rfds, nil, nil, timeout)
	if err != nil {
		if errors.Is(err, unix.EINTR) {
			return false, nil
		}
		return false, err
	}

	return n > 0 && rfds.IsSet(fd), nil
}
