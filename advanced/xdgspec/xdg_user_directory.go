package xdgspec

import (
	"runtime"

	"github.com/adrg/xdg"
	"github.com/devlights/try-golang/output"
)

// XdgUserDirectory は、XDG User Directory についてのサンプルです.
// [xdg](https://github.com/adrg/xdg) を利用して各値を取得しています.
//
// REFEFENCES::
//   - https://www.freedesktop.org/wiki/Software/xdg-user-dirs/
func XdgUserDirectory() error {
	// ------------------------------------------------------------
	// XDG User Directory について
	// XDG の各値は、xdg パッケージを利用すると簡単に利用できる
	// このパッケージは、Windows/MacOS/Unix に対応している
	//
	// XDG User Directory は、Documents,Downloads,Music,Desktopなどの
	// $HOME に配置されるユーザ共通のディレクトリのセットのこと
	//
	// xdg パッケージは、対象となる環境変数が設定されていれば、その値を返し
	// 設定されていない場合は、デフォルトの値を返すようになっている。
	// ------------------------------------------------------------
	output.Stdoutl("[OS]", runtime.GOOS)

	// XDG_DESKTOP_DIR
	// デスクトップの場所
	output.Stdoutl("XDG_DESKTOP_DIR", xdg.UserDirs.Desktop)

	// XDG_DOWNLOAD_DIR
	// ダウンロードの場所
	output.Stdoutl("XDG_DOWNLOAD_DIR", xdg.UserDirs.Download)

	// XDG_DOCUMENTS_DIR
	// ドキュメントの場所
	output.Stdoutl("XDG_DOCUMENTS_DIR", xdg.UserDirs.Documents)

	// XDGの標準規定にはないもの
	output.Stdoutl("Application dirs", xdg.ApplicationDirs)
	output.Stdoutl("Font dirs", xdg.FontDirs)

	return nil
}
