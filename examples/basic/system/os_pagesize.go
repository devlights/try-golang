package system

import (
	"fmt"
	"os"
)

// PageSize は、os.Getpagesize() のサンプルです.
func PageSize() error {
	// 仮想メモリアドレスから物理アドレス上のデータへのアクセスには
	// ページテーブルが利用される。メモリ管理のこの部分はGo言語からは
	// 直接触れることは出来ないが、ページサイズの情報を返すAPIは存在する.
	pagesize := os.Getpagesize()
	fmt.Printf("page size = %v\n", pagesize)

	return nil

	/*
	   $ task
	   task: Task "build" is up to date
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: os_getpagesize

	   [Name] "os_getpagesize"
	   page size = 4096


	   [Elapsed] 4.17µs
	*/

}
