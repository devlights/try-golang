package network

import (
	"net"
	"strconv"

	"github.com/devlights/gomy/output"
)

// SplitHostPort は、net.SplitHostPort のサンプルです。
//
// 文字列結合で 127.0.0.1:9999 という形を作っても問題はないが
// IPv6の場合は [::1]:9999 とする必要があるため、net.JoinHostPort() を
// 利用した方が間違いが少なくなる。
//
// 逆に分割したい場合は、同じように net.SplitHostPort を使った方が良い。
//
// # References
//   - https://blog.lufia.org/entry/2022/12/16/205728
//   - https://pkg.go.dev/net@go1.19.4#JoinHostPort
//   - https://www.geekpage.jp/blog/?id=2018-10-18-1
func SplitHostPort() error {
	var (
		port  = strconv.Itoa(9999)
		addrs = []string{
			net.JoinHostPort("", port),
			net.JoinHostPort("127.0.0.1", port),
			net.JoinHostPort("localhost", port),
			net.JoinHostPort("::1", port),
			net.JoinHostPort("fe80::a%eth0", port),
		}
	)

	for _, addr := range addrs {
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return err
		}

		output.Stdoutl("[Original     ]", addr)
		output.Stdoutf("[SplitHostPort]", "host=%v\tport=%v\n", host, port)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: network_split_host_port

	   [Name] "network_split_host_port"
	   [Original     ]      :9999
	   [SplitHostPort]      host=      port=9999
	   [Original     ]      127.0.0.1:9999
	   [SplitHostPort]      host=127.0.0.1     port=9999
	   [Original     ]      localhost:9999
	   [SplitHostPort]      host=localhost     port=9999
	   [Original     ]      [::1]:9999
	   [SplitHostPort]      host=::1   port=9999
	   [Original     ]      [fe80::a%eth0]:9999
	   [SplitHostPort]      host=fe80::a%eth0  port=9999


	   [Elapsed] 100.64µs
	*/

}
