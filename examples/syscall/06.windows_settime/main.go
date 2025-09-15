//go:build windows

package main

import (
	"log"
	"syscall"
	"time"
)

func main() {
	//
	// WindowsとLinuxでは atime,mtime,ctime の意味が異なる。
	//
	// Linux:
	//   - atime (Access Time)       - アクセス時間：ファイルが最後に読み取られた時刻
	//   - mtime (Modification Time) - 更新時間：ファイルの内容が最後に変更された時刻
	//   - ctime (Change Time)       - 変更時間：ファイルのメタデータ（権限、所有者、リンク数など）が最後に変更された時刻
	//
	// Windows:
	//   - atime (Access Time)   - アクセス時間 (Last Access Time)
	//   - mtime (Write Time)    - 更新時間 (Last Write Time)
	//   - ctime (Creation Time) - 作成時間 (Creation Time)
	//
	// Windowsの ctime (Creation Time) は変更可能であるが、Linuxの ctime (Change Time) は変更不可。
	// Windowsの場合に ３つの時間 (atime,mtime,ctime) を変更したい場合は syscall.SetFileTime() を用いる。
	//
	// osパッケージに存在する os.Chtime() では、WindowsのCreation Timeは変更できない。
	//
	//   func Chtimes(name string, atime time.Time, mtime time.Time) error
	//
	// syscallには SetFileTime 関数が存在する。
	//
	//   func SetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error)
	//
	// 第一引数の Handle も syscall の関数を用いて取得する.
	//
	log.SetFlags(0)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	const (
		FilePath = "./test.txt"
	)
	var (
		fname *uint16
		err   error
	)
	fname, err = syscall.UTF16PtrFromString(FilePath)
	if err != nil {
		return err
	}

	var (
		handle       syscall.Handle
		access       uint32 = syscall.FILE_WRITE_ATTRIBUTES
		mode         uint32 = syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE | syscall.FILE_SHARE_DELETE
		createMode   uint32 = syscall.OPEN_EXISTING
		attrs        uint32 = syscall.FILE_ATTRIBUTE_NORMAL
		templateFile int32  = 0
	)
	handle, err = syscall.CreateFile(fname, access, mode, nil, createMode, attrs, templateFile)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(handle)

	var (
		jst, _ = time.LoadLocation("Asia/Tokyo")
		a      = time.Date(2025, 1, 1, 0, 0, 0, 0, jst).UnixNano()
		m      = time.Date(2025, 1, 2, 0, 0, 0, 0, jst).UnixNano()
		c      = time.Date(2025, 1, 3, 0, 0, 0, 0, jst).UnixNano()
		atime  = syscall.NsecToFiletime(a)
		mtime  = syscall.NsecToFiletime(m)
		ctime  = syscall.NsecToFiletime(c)
	)
	err = syscall.SetFileTime(handle, &ctime, &atime, &mtime)
	if err != nil {
		return err
	}

	return nil
}
