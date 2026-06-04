package main

import (
	"encoding/hex"
	"log"
	"os"
	"unsafe"

	"github.com/goforj/godump"
)

const (
	DataSize = 10
)

func main() {
	log.SetFlags(0)

	// 1.byte配列をbyteスライスに普通に変換
	fn1()

	// 2.byte配列をunsafe.Sliceを使ってbyteスライスに変換
	fn2()

	// 3.byte配列をunsafe.Sliceを使ってuint16スライスに変換
	fn3()
}

func fn1() {
	//
	// 一番普通の変換
	//   [N]byte を [:] すればスライスになる
	//
	// fn2との比較:
	//   - こちらはGoの組み込み構文で型安全
	//   - cgoでC側のポインタを受け取る場合はfn2のパターンが必要になる
	//

	var (
		l     = log.New(os.Stdout, "[fn1] ", 0)
		data  [DataSize]byte
		slice []byte
	)
	for i := range len(data) {
		data[i] = uint8(i + 1)
	}
	slice = data[:]

	l.Printf("slice: addr=%p, len=%d, cap=%d", slice, len(slice), cap(slice))
	{
		// すでにlen==capとなっているので、このappendで拡張されてポインタが変わる
		slice = append(slice, 11)
	}
	l.Printf("slice: addr=%p, len=%d, cap=%d", slice, len(slice), cap(slice))

	godump.Dump(slice)
}

func fn2() {
	//
	// unsafe.Sliceを使ってキャストする
	// 普段は利用しないが、cgoではよく利用するパターン
	//
	// fn1との違い:
	//   - fn1はGoの配列変数に対してのみ使える
	//   - こちらはcgoで受け取った *C.uchar や *C.char など、
	//     Goの型システム外のポインタに対しても適用できる
	//
	//   例:
	//     // C側から受け取ったバイト列をGoのスライスとして扱う
	//     cPtr := C.get_buffer()
	//     slice := unsafe.Slice((*byte)(unsafe.Pointer(cPtr)), C.get_buffer_len())
	//

	var (
		l     = log.New(os.Stdout, "[fn2] ", 0)
		data  [DataSize]byte
		slice []byte
	)
	for i := range len(data) {
		data[i] = uint8(i + 1)
	}
	slice = unsafe.Slice((*uint8)(unsafe.Pointer(&data[0])), len(data))

	l.Printf("slice: addr=%p, len=%d, cap=%d", slice, len(slice), cap(slice))
	{
		// すでにlen==capとなっているので、このappendで拡張されてポインタが変わる
		slice = append(slice, 11)
	}
	l.Printf("slice: addr=%p, len=%d, cap=%d", slice, len(slice), cap(slice))

	godump.Dump(slice)
}

func fn3() {
	//
	// unsafe.Sliceを使ってベース型を変更する
	// 普段は利用しないが、cgoではよく利用するパターン
	//
	// NOTE: エンディアン依存
	//   byte列を uint16 として解釈する場合、結果はCPUのエンディアンに依存する。
	//   リトルエンディアン（x86/x64/ARM）では data[0] が uint16 の下位バイトになる。
	//
	//   例: data[0]=0x01, data[1]=0x02 の場合
	//     リトルエンディアン: uint16 = 0x0201 (= 513)
	//     ビッグエンディアン: uint16 = 0x0102 (= 258)
	//
	//   ネットワークバイトオーダー（ビッグエンディアン）のデータを扱う場合は
	//   encoding/binary パッケージの binary.BigEndian.Uint16() 等を使うこと。
	//
	// NOTE: アライメント
	//   uint16 は 2バイトアライメントが必要。
	//   今回は Go のスタック変数 [DataSize]byte を使っているため問題ないが、
	//   cgo で C 側から受け取ったポインタ（malloc等）の場合は
	//   アライメントが保証されているか事前に確認すること。
	//   アライメント違反は一部のCPUアーキテクチャ（SPARC, MIPS等）では SIGBUS を引き起こす。
	//   x86/x64 では SIGBUS は発生しないが、パフォーマンス低下の原因になる場合がある。
	//

	var (
		l     = log.New(os.Stdout, "[fn3] ", 0)
		data  [DataSize]byte
		slice []uint16
	)
	for i := range len(data) {
		data[i] = uint8(i + 1)
	}

	// byte配列を uint16 スライスとして見るため、要素数は len/2 になる
	// (uint16 は 2バイトなので、DataSize=10バイト → 5要素)
	slice = unsafe.Slice((*uint16)(unsafe.Pointer(&data[0])), len(data)/2)

	l.Printf("slice: addr=%p, len=%d, cap=%d", slice, len(slice), cap(slice))
	{
		// すでにlen==capとなっているので、このappendで拡張されてポインタが変わる
		slice = append(slice, 0x1211)
	}
	l.Printf("slice: addr=%p, len=%d, cap=%d", slice, len(slice), cap(slice))

	// 0番目の要素が 0x0201 (513) と表示される
	godump.Dump(slice)

	// 一応バイナリでも確認
	b := unsafe.Slice((*byte)(unsafe.Pointer(&slice[0])), len(slice)*2)
	log.Println(hex.Dump(b))
}
