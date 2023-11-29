package defers

import "fmt"

// DeferInLoop は、deferをループ内で利用したい場合のやり方についてのサンプルです。
//
// REFERENCES::
//   - https://mattn.kaoriya.net/software/lang/go/20151212021608.htm
//   - https://stackoverflow.com/questions/45617758/defer-in-the-loop-what-will-be-better
func DeferInLoop() error {
	// --------------------------------------------------------------
	// 上のURLに書かれているように、deferは内部でLIFOキューで管理されている
	// ので、ループ内で defer をそのまま書くと、どんどんキューに溜まってしまう.
	//
	// なので、ループ内で defer する場合は、匿名関数を用意して
	// 関数スコープを作り、その中で defer するようにする。
	//
	// 関数スコープが作られるので、ループが一回分終了するたびにスコープを
	// 抜け、その際に defer が実行される。
	//
	// ループ内でファイルをオープンしていく処理などを書いている場合に特に注意。
	// --------------------------------------------------------------

	// NGパターン
	bad()

	fmt.Println("---------------------------------------")

	// OKパターン
	good()

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: defer_in_loop

	   [Name] "defer_in_loop"
	   defer: 09
	   defer: 08
	   defer: 07
	   defer: 06
	   defer: 05
	   defer: 04
	   defer: 03
	   defer: 02
	   defer: 01
	   defer: 00
	   ---------------------------------------
	   defer: 00
	   defer: 01
	   defer: 02
	   defer: 03
	   defer: 04
	   defer: 05
	   defer: 06
	   defer: 07
	   defer: 08
	   defer: 09


	   [Elapsed] 139.21µs
	*/

}

func bad() {
	// ループ内で defer をそのまま使っているので
	// LIFOキューにループ回数分の defer が溜まった後
	// 関数を抜けるときに実行される。
	//
	// なので、この場合は
	//   defer: 09
	//   defer: 08
	//   ・
	//   ・
	//   defer: 00
	//
	// という風に、逆順で出力される
	for i := 0; i < 10; i++ {
		//noinspection GoDeferInLoop
		defer fmt.Printf("defer: %02d\n", i)
	}
}

func good() {
	// ループ内で defer をそのまま使わずに関数スコープを作り
	// defer を使っている。そのため、ループ一回毎に defer が
	// 実行される。
	//
	// なので、この場合は
	//   defer: 00
	//   defer: 01
	//   ・
	//   ・
	//   defer: 09
	//
	// という風に出力される
	for i := 0; i < 10; i++ {
		func(i int) {
			defer fmt.Printf("defer: %02d\n", i)
		}(i)
	}
}
