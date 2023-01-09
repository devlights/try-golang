package network

import (
	"net"

	"github.com/devlights/gomy/output"
)

// LookupPort は、 net.LookupPort() のサンプルです。
//
// ポートの範囲チェックも行ってくれるので外部からポート番号を受け取って
// 処理する場合は、 net.LookupPort() を使ったほうが良い。
//
// context.Context を利用したい場合は net.Resolver.LookupPort() を利用する。
// デフォルトの Resolver で良い場合は net.DefaultResolver.LookupPort() が使える。
//
// # REFERENCES
//   - https://blog.lufia.org/entry/2022/12/16/205728
//   - https://pkg.go.dev/net@go1.19.4#LookupPort
//   - https://www.infraexpert.com/study/tea5.htm
func LookupPort() error {
	const (
		protocol = "tcp"
	)

	var (
		services = []string{
			"ftp",
			"ssh",
			"telnet",
			"http",
			"pop3",
			"imap",
			"https",
			"8888",
			"-1",    // 範囲外
			"65536", // 範囲外
			"xdmcp", // xdmcpはUDPプロトコル (port=177)
		}
	)

	for _, service := range services {
		port, err := net.LookupPort(protocol, service)
		if err != nil {
			output.Stdoutf("[Err]", "service=%-5v\tErr=%v\n", service, err)
			continue
		}

		output.Stdoutf("[LookupPort]", "serivce=%v\tport=%v\n", service, port)
	}

	return nil
}
