package osop

import (
	"os"

	"github.com/devlights/gomy/output"
)

// OpenRoot は、os.OpenRoot についてのサンプルです。
//
// os.OpenRoot は、Go 1.24 で入った特定ディレクトリ配下に操作を閉じ込めるための API です 。
// 指定したディレクトリをrootとして扱い、その配下のパスのみを許可するようにします。
// つまり、外に(../)出れないようにします。
//
// 自分自身で利用する小さなツールなどでは利用しなくても問題ありませんが
// 外部から渡されてきたパスなど信頼できない可能性がある場合は、この関数を用いて閉じた操作をすると
// 安全性が高まります。(先に filepath.Clean を呼んで、パスのクリーニングもしておいたほうが良い)
//
// 戻り値は *os.Root となっており、ここから Open したり Read/Write したり出来ます。
//
// > Root may be used to only access files within a single directory tree.
// Methods on Root can only access files and directories beneath a root directory. If any component of a file name passed to a method of Root references a location outside the root, the method returns an error. File names may reference the directory itself (.).
// Methods on Root will follow symbolic links, but symbolic links may not reference a location outside the root. Symbolic links must not be absolute.
// Methods on Root do not prohibit traversal of filesystem boundaries, Linux bind mounts, /proc special files, or access to Unix device files.
// Methods on Root are safe to be used from multiple goroutines simultaneously.
// On most platforms, creating a Root opens a file descriptor or handle referencing the directory. If the directory is moved, methods on Root reference the original directory in its new location.
//
// > Root は、単一のディレクトリツリー内のファイルへのアクセスにのみ使用できます。
// Root のメソッドは、ルートディレクトリ以下のファイルとディレクトリにのみアクセスできます。Root のメソッドに渡されるファイル名の一部がルート以外の場所を参照している場合、メソッドはエラーを返します。ファイル名はディレクトリ自体 (.) を参照できます。
// Root のメソッドはシンボリックリンクをたどりますが、シンボリックリンクはルート以外の場所を参照できません。シンボリックリンクは絶対リンクであってはなりません。
// Root のメソッドは、ファイルシステム境界、Linux のバインドマウント、/proc 特殊ファイル、または Unix デバイスファイルへのアクセスを禁止しません。
// Root のメソッドは、複数のゴルーチンから同時に使用しても安全です。
// ほとんどのプラットフォームでは、Root を作成すると、ディレクトリを参照するファイルディスクリプタまたはハンドルが開かれます。ディレクトリが移動された場合、Root のメソッドは新しい場所にある元のディレクトリを参照します。
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.26.3#OpenRoot
//   - https://pkg.go.dev/os@go1.26.3#Root
func OpenRoot() error {
	var (
		cwd  string
		root *os.Root
		err  error
	)
	if cwd, err = os.Getwd(); err != nil {
		return err
	}
	if root, err = os.OpenRoot(cwd); err != nil {
		return err
	}
	defer root.Close()

	// ルート配下のファイルは操作可能
	var (
		pOk = "examples/basic/osop/openroot.go"
		fOk *os.File
	)
	if fOk, err = root.Open(pOk); err != nil {
		return err
	}
	defer fOk.Close()

	// ルートの外に出ようとするとエラー
	// (path escapes from parent)
	var (
		pNg = "../file.txt"
		fNg *os.File
	)
	if fNg, err = root.Open(pNg); err != nil {
		output.Stdoutf("[ERR]", "%v", err)
		return nil
	}
	defer fNg.Close()

	return nil
}
