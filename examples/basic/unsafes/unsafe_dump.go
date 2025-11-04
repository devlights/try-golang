package unsafes

import (
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"unsafe"
)

// Dump は、unsafeパッケージを使って構造体のメモリダンプを出力するサンプルです。
func Dump() error {
	type (
		// わざとパディングが入る構造体とする
		// size: 24bytes
		S1 struct {
			A uint8  // offset=0,  size=1, padding=3
			B uint32 // offset=4,  size=4, padding=0
			C uint64 // offset=8,  size=8, padding=0
			D uint16 // offset=16, size=2, padding=6
		}
	)

	// 構造体のメモリダンプ
	//   尚、実行すると以下のように
	//        00000000  ff 62 a6 00 ff ff ff ff  ff ff ff ff ff ff ff ff
	//        00000010  ff ff 1e 00 c0 00 00 00
	//   62 a6 のようなゴミが見える。これはGoはパディング部分をゼロ初期化しないため。
	//   Goの仕様では、構造体のフィールドは初期化されるが、アライメントのためのパディング領域は未定義となっている。
	var (
		s1        = S1{math.MaxUint8, math.MaxUint32, math.MaxUint64, math.MaxUint16}
		size      = unsafe.Sizeof(s1)
		ptr       = unsafe.Pointer(&s1)
		bytePtr   = (*byte)(ptr)
		byteSlice = (([]byte)(unsafe.Slice(bytePtr, size)))
	)
	hex.Dumper(os.Stdout).Write(byteSlice)
	fmt.Println("")

	// ゼロクリアしてから値を再設定して確認してみる
	// 以下のループはmemset(ptr, 0, size) と同じ感じ。
	for i := range int(size) {
		if i == 0 {
			*bytePtr = 0
		}

		ptr = unsafe.Add(ptr, 1)
		bytePtr = (*byte)(ptr)
		*bytePtr = 0
	}

	// 構造体の値を設定しなおし
	s1.A = math.MaxUint8
	s1.B = math.MaxUint32
	s1.C = math.MaxUint64
	s1.D = math.MaxUint16

	hex.Dumper(os.Stdout).Write(byteSlice)

	return nil

	/*
		$ task
		task: [build] go build -o "/home/dev/dev/github/try-golang/try-golang" .
		task: [run] ./try-golang -onetime

		ENTER EXAMPLE NAME: unsafe_dump

		[Name] "unsafe_dump"
		00000000  ff 00 00 00 ff ff ff ff  ff ff ff ff ff ff ff ff  |................|
		00000010  ff ff 18 00 c0 00 00 00
		00000000  ff 00 00 00 ff ff ff ff  ff ff ff ff ff ff ff ff  |................|
		00000010  ff ff 00 00 00 00 00 00

		[Elapsed] 88.436µs
	*/
}
