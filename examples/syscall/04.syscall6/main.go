//go:build unix

package main

import (
	"log"
	"unsafe"

	"golang.org/x/sys/unix"
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// unix.Syscall6()を使って mmap(2) を呼び出し
	//

	// ファイルディスクリプタを取得
	var (
		fileName = "/etc/hosts"
		fd       int
		err      error
	)
	fd, err = unix.Open(fileName, unix.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer unix.Close(fd)

	// ファイルサイズを取得
	var (
		stat  unix.Stat_t
		fsize int64
	)
	err = unix.Fstat(fd, &stat)
	if err != nil {
		return err
	}

	fsize = stat.Size

	// mmap(2)を利用してメモリにマッピング
	//   引数については man 2 mmap で調べられる
	//
	// そもそも、unixパッケージには
	//   - unix.Mmap()
	//   - unix.Munmap()
	// が存在するので、それを利用するのが普通です。
	//
	// 本サンプルは、unix.Syscall6()を試す為の実装です。
	var (
		trap   = uintptr(unix.SYS_MMAP)
		addr   = uintptr(0)               // アドレスヒント(0は自動割当)
		length = uintptr(fsize)           // ファイルサイズ
		prot   = uintptr(unix.PROT_READ)  // 読み取り専用
		flags  = uintptr(unix.MAP_SHARED) // マッピングフラグ（共有）
		fdptr  = uintptr(fd)              // ファイルディスクリプタ
		offset = uintptr(0)               // オフセット
		errno  unix.Errno

		unmap = func(addr uintptr) {
			unix.Syscall(unix.SYS_MUNMAP, addr, length, uintptr(0))
		}
	)
	addr, _, errno = unix.Syscall6(trap, addr, length, prot, flags, fdptr, offset)
	if errno != 0 {
		err = errno
		return err
	}
	defer unmap(addr)

	// 先頭 Nバイト 分を見る
	const (
		dispSize = 50
	)
	var (
		// VSCodeなどで作業していると「possible misuse of unsafe.Pointer」という警告が出るが
		// mmapの戻り値はメモリアドレスを表すuintptr値なので、以下のunsafe.Pointerへの変換は安全です。
		ptr  = (*byte)(unsafe.Pointer(addr))
		data = unsafe.Slice(ptr, int(fsize))
	)
	if fsize >= dispSize {
		log.Println(string(data[:dispSize]))
	} else {
		log.Println(string(data))
	}

	return nil
}
