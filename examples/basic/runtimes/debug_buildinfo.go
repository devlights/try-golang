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
}
