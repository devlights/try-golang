package main

/*
extern void c_func2(const char *s, size_t n);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export go_func
func go_func(s *C.char, n C.size_t) {
	var (
		sPtr   = unsafe.Pointer(s)
		cSlice = unsafe.Slice((*byte)(sPtr), n) // cSliceはC側のスタック変数を指している
	)
	fmt.Printf("[Go] %s\n", cSlice)

	// 何らかの変換を行う（例としてデータをリバース）
	//
	// 注意点として、cSliceはC側のスタック変数をそのまま指しているため
	// これを直接変更すると、C側のスタックメモリを書き換えてしまうことになる。
	// 必ず、コピーを取ってから変更処理は行うこと。
	//
	// また、C.GoBytes(), C.CString() を利用せずに直接C側のデータを扱っているので
	// cSliceの中は最後に終端文字が入った状態となっている。
	// この状態でそのままスライスをリバースすると \0 が先頭に来ることになるので除去してから処理する。
	var (
		goSlice []byte   // Go側で扱うスライス
		dataLen = int(n) // 実データのサイズ
	)
	// NULL終端文字がある場合は減算して実データサイズとする
	if dataLen > 0 && cSlice[dataLen-1] == 0 {
		dataLen--
	}

	// 実データ分をコピー
	goSlice = make([]byte, dataLen)
	copy(goSlice, cSlice[:dataLen])

	// リバース
	for i, j := 0, len(goSlice)-1; i < j; i, j = i+1, j-1 {
		goSlice[i], goSlice[j] = goSlice[j], goSlice[i]
	}

	// 終端追加
	goSlice = append(goSlice, 0)

	// C側の関数に渡すための準備
	var (
		bytePtr = unsafe.SliceData(goSlice)          // *byteに変換し
		charPtr = (*C.char)(unsafe.Pointer(bytePtr)) // そこから (char *) に変換
		charLen = C.size_t(len(goSlice))             // サイズはスライスからそのまま取得 ([]byteの場合はこれでOK)
	)
	C.c_func2(charPtr, charLen)
}
