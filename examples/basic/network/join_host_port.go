package network

import (
	"net"
	"strconv"

	"github.com/devlights/gomy/output"
)

// JoinHostPort は、net.JoinHostPort のサンプルです。
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
func JoinHostPort() error {
	var (
		hosts = []string{
			"",
			"127.0.0.1",
			"localhost",
			"::1",
		}
		port = strconv.Itoa(9999)
	)

	for _, host := range hosts {
		addr := net.JoinHostPort(host, port)
		output.Stdoutl("[JoinHostPort]", addr)
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: network_join_host_port

	   [Name] "network_join_host_port"
	   [JoinHostPort]       :9999
	   [JoinHostPort]       127.0.0.1:9999
	   [JoinHostPort]       localhost:9999
	   [JoinHostPort]       [::1]:9999


	   [Elapsed] 18.86µs
	*/

}
