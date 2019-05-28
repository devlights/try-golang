package slice_

import "fmt"

// スライスについてのサンプル
// 空のスライスとnilなスライスの違い
func Slice04() error {
	// REFERENCES:: http://bit.ly/2I3cwuD
	slice1 := []int{1, 2, 3, 4, 5}
	printSliceInfo("slice1", slice1)

	// 全要素を削除 (lenもcapも0にする)
	slice1 = nil
	printSliceInfo("slice1", slice1) // => [] len: 0 cap: 0

	// 空のスライスにする
	// この場合メモリの開放は行われない
	slice2 := []int{1, 2, 3, 4, 5}
	slice2 = slice2[:0]
	printSliceInfo("slice2", slice2) // => [] len: 0 cap: 5

	// cap はそのままで残るため、メモリの開放は行われていない
	// なので、再度拡張してやると元々のデータがまた見える
	slice2 = slice2[:3]
	printSliceInfo("slice2", slice2)

	return nil
}

func printSliceInfo(prefix string, slice []int) {
	fmt.Printf(
		"[%slice] val: %v len: %d cap: %d\n",
		prefix,
		slice,
		len(slice),
		cap(slice))
}
