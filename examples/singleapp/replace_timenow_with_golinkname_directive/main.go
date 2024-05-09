package main

import (
	"log"
	"time"
	_ "unsafe"
)

// fixedNow は、固定の時間を返す time.Now() のモックです。
//
// 標準ライブラリの time.Now() を置き換えるために
// linknameコンパイラディレクティブを付与しています。
//
// linknameコンパイラディレクティブを利用するためには
// unsafeパッケージをインポート（明示的、暗黙的問わず）する必要があります。
//
// linknameは、標準ライブラリ内で利用されており、例えば time.Sleep() も
// 実際のソースコードを見ると以下の宣言となっており、宣言のみで実装がありません。
//
// # $(go env GOROOT)/src/time/sleep.go
//
//	package time
//
//	func Sleep(d Duration)
//
// 実体は、runtime.timeSleep() となっています。
//
// # $(go env GOROOT)/src/runtime/time.go
//
//	package runtime
//
//	//go:linkname timeSleep time.Sleep
//	func timeSleep(ns int64) { ... }
//
//go:linkname fixedNow time.Now
func fixedNow() time.Time {
	return time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local)
}

func init() {
	log.SetFlags(0)
}

func main() {
	// linknameコンパイラディレクティブによって time.Now が置き換えられる
	log.Println(time.Now())
}
