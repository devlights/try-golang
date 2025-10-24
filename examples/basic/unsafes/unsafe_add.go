package unsafes

import (
	"fmt"
	"unsafe"
)

// Add は、unsafe.Add関数を利用してポインタ演算するサンプルです。
//
// unsafe.Addは、Go 1.17で追加された関数。
// それまでは unsafe.Pointer(uintptr(ptr) + uintptr(len)) という形で行っていたポインタ演算を
// 内部で行ってくれるヘルパー関数。
//
// これを利用することで Go でも ポインタ演算 出来るようになる。
// ただし、unsafeパッケージを利用する時点でGoの持つ安全性を無くすことに注意が必要。
//
// REFERENCES:
//   - https://pkg.go.dev/unsafe@go1.25.3#Add
func Add() error {
	var (
		// 0x01020304 をリトルエンディアン環境で格納すると
		// メモリ上は [04][03][02][01] の順に配置される
		value      int = 0x01020304
		ptr            = unsafe.Pointer(&value)
		bytePtr        = (*byte)(ptr)
		sizeOfByte     = unsafe.Sizeof((byte)(0))
	)

	//
	// --- byte単位でのアクセス (1バイトずつ読み取り) ---
	//

	// 先頭バイト（最下位バイト）を表示
	// リトルエンディアンなので 0x04 が格納されている
	fmt.Printf("1バイト目: 0x%02X\n", *bytePtr)

	// アドレスを (1 * sizeof(byte)) 分進める (Go 1.17までのやり方)
	// uintptrに変換してからオフセット加算し、再度unsafe.Pointerに戻す
	//
	// (重要)
	// この方法で処理する場合は、ポインタ演算部分は1ステートメントで行うことが必須。
	// uintptrは単なる整数型であり、GCはuintptrが指すオブジェクトを追跡しない。
	// そのため、uintptr型の値を変数に保持したまま別の処理を挟むと、
	// その間にGCが実行され、元のオブジェクトが移動または解放される可能性がある。
	//
	// NG:
	//
	//	tmp := uintptr(ptr) + offset  // この時点でGCが入る可能性
	//	newPtr := unsafe.Pointer(tmp) // tmpが指すアドレスは既に無効かもしれない
	//
	// OK:
	//
	//	newPtr := unsafe.Pointer(uintptr(ptr) + offset)  // 1ステートメントで完結
	//
	ptr = unsafe.Pointer(uintptr(ptr) + sizeOfByte)
	bytePtr = (*byte)(ptr)

	// 2バイト目を表示 (0x03)
	fmt.Printf("2バイト目: 0x%02X\n", *bytePtr)

	// アドレスを (2 * sizeof(byte)) 分進める (Go 1.17以降のやり方)
	//
	// unsafe.Add を使うことで、より簡潔にポインタ演算が可能
	// (unsafe.Addは内部で1ステートメントとして処理される)
	//
	// 元の位置から2バイト進めるので、4バイト目(0x01)にアクセス
	ptr = unsafe.Add(ptr, 2*sizeOfByte)
	bytePtr = (*byte)(ptr)

	// 4バイト目（最上位バイト）を表示 (0x01)
	fmt.Printf("4バイト目: 0x%02X\n", *bytePtr)

	//
	// --- int16単位でのアクセス (2バイトずつ読み取り) ---
	// キャスト先を *byte から *int16 へ変更
	//

	var (
		ptr2        = unsafe.Pointer(&value) // 先頭アドレスを再取得
		shortPtr    = (*int16)(ptr2)
		sizeOfShort = unsafe.Sizeof((int16(0)))
	)

	// 最初の2バイト [04][03] をint16として読み取り
	// リトルエンディアンなので 0x0304 と解釈される
	fmt.Printf("short: 0x%04X\n", *shortPtr)

	// 2バイト(1 * sizeof(short))進めて、次の2バイト [02][01] をint16として読み取り
	// リトルエンディアンなので 0x0102 と解釈される
	shortPtr = (*int16)(unsafe.Add(ptr2, sizeOfShort))
	fmt.Printf("short: 0x%04X\n", *shortPtr)

	return nil

	/*
	   $ task
	   task: [build] go build -o "/home/dev/dev/github/try-golang/try-golang" .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: unsafe_add

	   [Name] "unsafe_add"
	   1バイト目: 0x04
	   2バイト目: 0x03
	   4バイト目: 0x01
	   short: 0x0304
	   short: 0x0102

	   [Elapsed] 12.542µs
	*/
}
