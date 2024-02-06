package runtimes

import (
	"errors"
	"runtime/debug"

	"github.com/devlights/gomy/output"
)

// DebugBuildInfo は、runtime/debug.ReadBuildInfo() のサンプルです.
//
// # REFERENCES
//   - https://pkg.go.dev/runtime/debug@go1.21.5#ReadBuildInfo
//   - https://pkg.go.dev/runtime/debug@go1.21.5#BuildSetting
func DebugBuildInfo() error {
	//
	// ReadBuildInfo()は、ReadBuildInfo は、実行中のバイナリに埋め込まれたビルド情報を返す.
	// 中にビルド時のGoのバージョンや依存関係、VCSのリビジョンなどが設定されている.
	//

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return errors.New("error: call debug.ReadBuildInfo()")
	}

	output.Stdoutl("[BuildInfo]", info)
	output.StdoutHr()

	output.Stdoutl("[Go Version]", info.GoVersion)
	output.Stdoutl("[Path]", info.Path)

	const revisionKey = "vcs.revision"
	for _, s := range info.Settings {
		if s.Key == revisionKey {
			output.Stdoutl("[revision]", s.Value)
		}
	}

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: runtime_debug_buildinfo

	   [Name] "runtime_debug_buildinfo"
	   [BuildInfo]          go go1.21.6
	   path    github.com/devlights/try-golang
	   mod     github.com/devlights/try-golang (devel)
	   dep     github.com/devlights/gomy       v0.6.0  h1:7BT8bSxr+ZeNkgEYNufuM2rSc6kIoN6g2FSZvrcT9zw=
	   dep     github.com/pelletier/go-toml/v2 v2.1.1  h1:LWAJwfNvjQZCFIDKWYQaM62NcYeYViCmWIwmOStowAI=
	   dep     github.com/shopspring/decimal   v1.3.1  h1:2Usl1nmF/WZucqkFZhnfFYxxxu8LG21F6nPQBE5gKV8=
	   dep     golang.org/x/crypto     v0.18.0 h1:PGVlW0xEltQnzFZ55hkuX5+KLyrMYhHld1YHO4AKcdc=
	   dep     golang.org/x/exp        v0.0.0-20240119083558-1b970713d09a      h1:Q8/wZp0KX97QFTc2ywcOE0YRjZPVIx+MXInMzdvQqcA=
	   dep     golang.org/x/sync       v0.6.0  h1:5BMeUDZ7vkXGfEr1x9B4bRcTH4lpkTkpdh0T/J+qjbQ=
	   dep     golang.org/x/sys        v0.16.0 h1:xWw16ngr6ZMtmxDyKyIgsE93KNKz5HKmMa3b8ALHidU=
	   dep     golang.org/x/term       v0.16.0 h1:m+B6fahuftsE9qjo0VWp2FW0mB3MTJvR0BaMQrq0pmE=
	   dep     golang.org/x/text       v0.14.0 h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=
	   dep     gopkg.in/natefinch/lumberjack.v2        v2.2.1  h1:bBRl1b0OH9s/DuPhuXpNl+VtCaJXFZ5/uEFST95x9zc=
	   dep     gopkg.in/yaml.v3        v3.0.1  h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
	   build   -buildmode=exe
	   build   -compiler=gc
	   build   CGO_ENABLED=1
	   build   CGO_CFLAGS=
	   build   CGO_CPPFLAGS=
	   build   CGO_CXXFLAGS=
	   build   CGO_LDFLAGS=
	   build   GOARCH=amd64
	   build   GOOS=linux
	   build   GOAMD64=v1
	   build   vcs=git
	   build   vcs.revision=a71b8c37618dedfcb77af1728fbe8359a37704e4
	   build   vcs.time=2024-01-31T05:28:38Z
	   build   vcs.modified=true

	   --------------------------------------------------
	   [Go Version]         go1.21.6
	   [Path]               github.com/devlights/try-golang
	   [revision]           a71b8c37618dedfcb77af1728fbe8359a37704e4

	   [Elapsed] 75.61µs
	*/
}
