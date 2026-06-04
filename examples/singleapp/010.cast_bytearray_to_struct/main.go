package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"unsafe"
)

type ST struct {
	V1    int32
	V2    uint32
	V3    [10]byte
	Dummy [6]byte
}

func (me *ST) String() string {
	return fmt.Sprintf("V1=%d, V2=%d, V3=%s, Dummy=%v\n", me.V1, me.V2, string(me.V3[:]), me.Dummy)
}

func main() {
	//
	// バイト列を作成
	//
	var (
		buf bytes.Buffer
		bin []byte
	)

	binary.Write(&buf, binary.LittleEndian, int32(127))
	binary.Write(&buf, binary.LittleEndian, uint32(255))
	buf.WriteString("helloworld")
	buf.Write(make([]byte, 6))

	bin = buf.Bytes()
	fmt.Println(hex.Dump(bin))

	//
	// 構造体にキャスト
	//
	var (
		ptr unsafe.Pointer
		st  *ST
	)

	ptr = unsafe.Pointer(&bin[0])
	st = (*ST)(ptr)

	fmt.Printf("%v\n", st)

	/*
	   00000000  7f 00 00 00 ff 00 00 00  68 65 6c 6c 6f 77 6f 72  |........hellowor|
	   00000010  6c 64 00 00 00 00 00 00                           |ld......|

	   V1=127, V2=255, V3=helloworld, Dummy=[0 0 0 0 0 0]
	*/

}
