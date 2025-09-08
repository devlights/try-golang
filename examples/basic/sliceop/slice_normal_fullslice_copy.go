package sliceop

import "fmt"

// NormalFullsliceCopy は、通常スライス、フルスライス式、copyビルドイン関数を利用した場合のサンプルです。
func NormalFullsliceCopy() error {
	// 通常のスライス操作（メモリ共有あり）
	{
		s := make([]int, 0, 5)
		s = append(s, 1, 2, 3, 4, 5)

		s2 := s[:3]
		s2 = append(s2, 100)

		s3 := s[3:]
		s3 = append(s3, 200)

		fmt.Println(cap(s), cap(s2), cap(s3)) // 5 5 4
		fmt.Println(s, s2, s3)                // [1 2 3 100 5] [1 2 3 100] [100 5 200]
	}
	// フルスライス式（サブスライス作成時に容量制限しメモリ共有を防ぐ）
	{
		s := make([]int, 0, 5)
		s = append(s, 1, 2, 3, 4, 5)

		s2 := s[:3:3]
		s2 = append(s2, 100)

		s3 := s[3:len(s):len(s)]
		s3 = append(s3, 200)

		fmt.Println(cap(s), cap(s2), cap(s3)) // 5 6 4
		fmt.Println(s, s2, s3)                // [1 2 3 4 5] [1 2 3 100] [4 5 200]
	}
	// copyビルドイン関数の利用
	{
		s := make([]int, 0, 5)
		s = append(s, 1, 2, 3, 4, 5)

		s2 := make([]int, 3)
		copy(s2, s[:3])
		s2 = append(s2, 100)

		s3 := make([]int, 2)
		copy(s3, s[3:])
		s3 = append(s3, 200)

		fmt.Println(cap(s), cap(s2), cap(s3)) // 5 6 4
		fmt.Println(s, s2, s3)                // [1 2 3 4 5] [1 2 3 100] [4 5 200]
	}

	return nil
}
