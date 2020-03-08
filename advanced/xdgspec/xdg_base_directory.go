package xdgspec

import (
	"runtime"

	"github.com/adrg/xdg"
	"github.com/devlights/try-golang/lib/output"
)

// XdgBaseDirectory は、XDG Base Directory についてのサンプルです.
// [xdg](https://github.com/adrg/xdg) を利用して各値を取得しています.
//
// REFEFENCES::
//   - https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html
//   - https://kledgeb.blogspot.com/2013/04/ubuntu-10-xdg-base-directory.html
func XdgBaseDirectory() error {
	// ------------------------------------------------------------
	// XDG Base Directory について
	// XDG の各値は、xdg パッケージを利用すると簡単に利用できる
	// このパッケージは、Windows/MacOS/Unix に対応している
	//
	// xdg パッケージは、対象となる環境変数が設定されていれば、その値を返し
	// 設定されていない場合は、デフォルトの値を返すようになっている。
	// ------------------------------------------------------------
	output.Stdoutl("[OS]", runtime.GOOS)

	// XDG_DATA_HOME
	// ユーザ個別のデータファイルが書き込まれる場所 (ユーザディレクトリ)
	output.Stdoutl("XDG_DATA_HOME", xdg.DataHome)

	// XDG_CONFIG_HOME
	// ユーザ個別の設定が書き込まれる場所 (ユーザディレクトリ)
	output.Stdoutl("XDG_CONFIG_HOME", xdg.ConfigHome)

	// XDG_CACHE_HOME
	// ユーザ毎のキャッシュデータの置き場 (ユーザディレクトリ)
	output.Stdoutl("XDG_CACHE_HOME", xdg.CacheHome)

	// XDG_RUNTIME_DIR
	// ユーザ毎の実行時ファイルやその他のファイルを置くべき場所 (ユーザディレクトリ)
	output.Stdoutl("XDG_RUNTIME_DIR", xdg.RuntimeDir)

	// XDG_DATA_DIRS
	// データファイルを検索する際のサーチパス (システムディレクトリ)
	//
	// (補足) アプリのデータを検索する場合は $XDG_DATA_HOME:$XDG_DATA_DIRS の順で行い
	// ユーザ毎のデータを優先させる
	output.Stdoutl("XDG_DATA_DIRS", xdg.DataDirs)

	// XDG_CONFIG_DIRS
	// 設定ファイルを検索する際のサーチパス (システムディレクトリ)
	//
	// (補足) アプリの設定を検索する場合は $XDG_CONFIG_HOME:$XDG_CONFIG_DIRS の順で行い
	// ユーザ毎の設定を優先させる
	output.Stdoutl("XDG_CONFIG_DIRS", xdg.ConfigDirs)

	return nil
}
