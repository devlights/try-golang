package osop

import (
	"cmp"
	"fmt"
	"os"
)

// SameFile は、os.SameFile のサンプルです。
//
// 2つの os.FileInfo が同一であるかを判定
// ファイル名や内容が一致するかでは無く
// 「ファイルシステム上で同じファイルを参照しているか」を判定する。
//
// UNIX系OSでは、以下の２点が比較される。
//
//   - デバイス番号
//   - inode番号
//
// ２つで判定するのは
// inode番号だけでは、異なるファイルシステム上に同じ番号が存在する可能性があるため。
//
// # REFERENCES
//   - https://pkg.go.dev/os@go1.27.2#SameFile
func SameFile() error {
	const (
		original = "original.txt"
		link     = "symlink.txt"
	)
	var (
		drop    = func(err error) {}
		prepare = func() func() {
			drop(os.WriteFile(original, []byte("helloworld"), 0640))
			drop(os.Symlink(original, link))
			return func() {
				drop(os.Remove(link))
				drop(os.Remove(original))
			}
		}
	)
	defer prepare()()

	var (
		fi1, err1 = os.Stat(original)
		fi2, err2 = os.Stat(link)
		fi3, err3 = os.Lstat(original)
		fi4, err4 = os.Lstat(link) // os.Statはリンクを追うが、os.Lstatは追わない

		err error
	)
	if err = cmp.Or(err1, err2, err3, err4); err != nil {
		return err
	}

	fmt.Printf("os.samefile(original-stat, link-stat)   ==> %v\n", os.SameFile(fi1, fi2))
	fmt.Printf("os.samefile(original-stat, link-lstat)  ==> %v\n", os.SameFile(fi1, fi4))
	fmt.Printf("os.samefile(original-lstat, link-stat)  ==> %v\n", os.SameFile(fi3, fi2))
	fmt.Printf("os.samefile(original-lstat, link-lstat) ==> %v\n", os.SameFile(fi3, fi4))

	return nil
}
