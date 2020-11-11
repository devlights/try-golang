package network

import (
	"net"

	"github.com/devlights/gomy/output"
)

// IpAddressParse -- net.ParseIP() の サンプルです.
func IpAddressParse() error {
	// ---------------------------------------------------------------
	// net.ParseIP() で 文字列 から IP を取得できる
	// 解析に失敗した場合は、 nil が返る
	// ---------------------------------------------------------------
	for _, s := range []string{"127.0.0.1", "127.0.01", "invalid"} {
		ip := net.ParseIP(s)
		if ip == nil {
			output.Stdoutl("[NG]", s)
		} else {
			output.Stdoutl("[OK]", ip)
		}
	}

	return nil
}
