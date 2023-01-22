package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"syscall"
	"time"
)

func main() {
	var (
		proto   = "tcp"
		addr    = net.JoinHostPort("localhost", "33333")
		backlog = 1
	)

	listener, err := listen(addr, backlog)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	// タイムアウトさせるために、わざと詰まらせる
	net.Dial(proto, addr)
	net.Dial(proto, addr)

	//
	// Contextを指定したい場合は net.Dialer or net.DefaultResolver().Dial を使う
	//
	var (
		connLog  = log.New(os.Stdout, "[With Context] ", log.Ltime)
		ctx, cxl = context.WithTimeout(context.Background(), 3*time.Second)
		dialer   = net.Dialer{}
	)
	defer cxl()

	connLog.Println("start")
	{
		_, err = dialer.DialContext(ctx, proto, addr)
		if err != nil {
			connLog.Println(err)
		}
	}
	connLog.Println("done")

	//
	// net.DialTimeout を使うという方法もある
	//
	connLog = log.New(os.Stdout, "[DialTimeout] ", log.Ltime)

	connLog.Println("start")
	{
		_, err = net.DialTimeout(proto, addr, 2*time.Second)
		if err != nil {
			connLog.Println(err)
		}
	}
	connLog.Println("done")
}

// listen は、指定された情報で net.Listener を生成して返します.
//
// # REFERENCES
//   - https://github.com/golang/go/issues/41470
//   - https://github.com/valyala/tcplisten/blob/master/tcplisten.go
//   - https://stackoverflow.com/a/49593356
func listen(addr string, backLog int) (net.Listener, error) {
	// make tcp addr
	var (
		tcpAddr *net.TCPAddr
		err     error
	)

	tcpAddr, err = net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		return nil, err
	}

	// make socket addr
	var (
		sockAddr syscall.SockaddrInet4
	)

	sockAddr.Port = tcpAddr.Port
	copy(sockAddr.Addr[:], tcpAddr.IP.To4())

	// make socket file descriptor
	var (
		sockFd     int
		sockDomain = syscall.AF_INET
		sockType   = syscall.SOCK_STREAM | syscall.SOCK_NONBLOCK | syscall.SOCK_CLOEXEC
		sockProto  = syscall.IPPROTO_TCP
	)

	sockFd, err = syscall.Socket(sockDomain, sockType, sockProto)
	if err != nil {
		return nil, err
	}

	// bind
	err = syscall.Bind(sockFd, &sockAddr)
	if err != nil {
		return nil, err
	}

	// listen
	err = syscall.Listen(sockFd, backLog)
	if err != nil {
		return nil, err
	}

	// make net.Listener
	var (
		fname    = fmt.Sprintf("backlog.%d.%s.%s", os.Getpid(), "tcp4", addr)
		file     = os.NewFile(uintptr(sockFd), fname)
		listener net.Listener
	)
	defer file.Close()

	listener, err = net.FileListener(file)
	if err != nil {
		return nil, err
	}

	return listener, nil
}
