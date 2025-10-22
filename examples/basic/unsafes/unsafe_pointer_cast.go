package unsafes

import (
	"fmt"
	"unsafe"
)

// PointerCast は、unsafeパッケージを用いてポインタを任意の型にキャストするサンプルです。
//
// unsafe.Pointer は、C言語でいう (void *) と同じものとなる。
// C言語において、(void *) は何にでも成れるのと同様に unsafe.Pointer はGoでもどの型にもキャスト出来る。
// ただし、unsafeパッケージを利用する時点でGoの持つ安全性を無くすことに注意が必要。
//
// REFERENCES:
//   - https://pkg.go.dev/unsafe@go1.25.3#Pointer
func PointerCast() error {
	var (
		value int = 0x01020304
		ptr   unsafe.Pointer
		cast1 *byte
		cast2 *[4]byte
	)
	// 元の値をunsafe.Pointerにする
	ptr = unsafe.Pointer(&value)

	// C言語でいう (char *)ptr; のような変換
	cast1 = (*byte)(ptr)

	// 明示的な配列のポインタへの変換
	cast2 = (*[4]byte)(ptr)

	// 値を確認 (cast1)
	// *cast1 は最初の1バイトのみを参照する (リトルエンディアン環境では 0x04)
	// 後続バイトにアクセスするには unsafe.Add() または uintptrを使ったポインタ演算 などを使用する必要がある
	fmt.Printf("cast1: 1バイト目: 0x%02X\n", *cast1)
	fmt.Println("---------------------------------")

	// 値を確認 (cast2)
	// *[4]byte 型なので配列として全4バイトに直接アクセス可能
	// メモリレイアウトはエンディアンに依存する（リトルエンディアンでは逆順）
	for i, v := range cast2 {
		fmt.Printf("cast2: %dバイト目: 0x%02X\n", i, v)
	}

	return nil

	/*
	   $ task
	   task: [build] go build -o "/home/dev/dev/github/try-golang/try-golang" .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: unsafe_pointer

	   [Name] "unsafe_pointer_cast"
	   cast1: 1バイト目: 0x04
	   ---------------------------------
	   cast2: 0バイト目: 0x04
	   cast2: 1バイト目: 0x03
	   cast2: 2バイト目: 0x02
	   cast2: 3バイト目: 0x01


	   [Elapsed] 42.422µs
	*/

}
