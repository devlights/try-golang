//go:build linux

package main

import (
	"encoding/binary"
	"log"
	"net"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	log.SetFlags(0)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// TCPリスナーの起動までを Syscall() で実装
	// なお、利用する関数は SyscallNoError (エラーを返さないバージョン) を意図的に利用
	//
	const (
		zero = uintptr(0) // 不要な引数値を表す
	)

	// socket(2)
	var (
		domain   = uintptr(unix.AF_INET)
		sockType = uintptr(unix.SOCK_STREAM)
		protocol = uintptr(0)
		sFd, _   = unix.SyscallNoError(unix.SYS_SOCKET, domain, sockType, protocol)
	)
	defer func(fd uintptr) {
		_, _ = unix.SyscallNoError(unix.SYS_CLOSE, fd, zero, zero)
	}(sFd)

	// setsockopt(2) (SO_REUSEADDR)
	var (
		level     = uintptr(unix.SOL_SOCKET)
		optname   = uintptr(unix.SO_REUSEADDR)
		optval    = 1
		optvalPtr = uintptr(unsafe.Pointer(&optval))
		optlen    = uintptr(unsafe.Sizeof(optval))
	)
	_, _, _ = unix.Syscall6(unix.SYS_SETSOCKOPT, sFd, level, optname, optvalPtr, optlen, zero)

	// ソケットアドレス生成
	//   アドレスを表現する構造体として unix.SockaddrInet4 と unix.RawSockaddrInet4 の２つがあるが
	//   unix.Syscall()を利用して直接システムコールを呼び出す場合は、カーネルが期待するメモリレイアウトを
	//   そのまま表現する unix.RawSockaddrInet4 を利用する。(要はCの構造体と同じ形の方を使う)
	//   unix.RawSockaddrInet4の方は、ネットワークバイトオーダーで値を保持する。
	var (
		sAddr    unix.RawSockaddrInet4
		sAddrPtr uintptr
		sAddrLen uintptr
	)
	sAddr = unix.RawSockaddrInet4{
		Family: unix.AF_INET,
		Port:   htons(8888),
		Addr:   [4]byte{127, 0, 0, 1},
	}
	sAddrPtr = uintptr(unsafe.Pointer(&sAddr))
	sAddrLen = uintptr(unix.SizeofSockaddrInet4)

	// bind(2)
	_, _ = unix.SyscallNoError(unix.SYS_BIND, sFd, sAddrPtr, sAddrLen)

	// listen(2)
	_, _ = unix.SyscallNoError(unix.SYS_LISTEN, sFd, uintptr(5), zero)

	for {
		// accept(2)
		var (
			cAddr       unix.RawSockaddrInet4
			cAddrPtr           = uintptr(unsafe.Pointer(&cAddr))
			cAddrLen    uint32 = unix.SizeofSockaddrInet4
			cAddrLenPtr        = uintptr(unsafe.Pointer(&cAddrLen))
			cFd, _             = unix.SyscallNoError(unix.SYS_ACCEPT, sFd, cAddrPtr, cAddrLenPtr)
		)
		log.Printf("[accept] EP: %v:%d", net.IP(cAddr.Addr[:]), ntohs(cAddr.Port))

		go func(fd uintptr) {
			time.Sleep(1 * time.Second)
			_, _ = unix.SyscallNoError(unix.SYS_CLOSE, fd, zero, zero)
		}(cFd)
	}
}

// ホストバイトオーダーからネットワークバイトオーダーに変換（Host to Network Short）
func htons(port uint16) uint16 {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, port)
	return binary.BigEndian.Uint16(bytes)
}

// ネットワークバイトオーダーからホストバイトオーダーに変換（Network to Host Short）
func ntohs(port uint16) uint16 {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, port)
	return binary.LittleEndian.Uint16(bytes)
}
