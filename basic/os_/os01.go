package os_

import (
	"fmt"
	"os"
)

// os.Getpagesize() のサンプル
// REFERENCES:: http://bit.ly/2R1izE5
func Os01() error {
	// 仮想メモリアドレスから物理アドレス上のデータへのアクセスには
	// ページテーブルが利用される。メモリ管理のこの部分はGo言語からは
	// 直接触れることは出来ないが、ページサイズの情報を返すAPIは存在する.
	pagesize := os.Getpagesize()
	fmt.Printf("page size = %v\n", pagesize)

	return nil
}
