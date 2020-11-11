package slices

import (
	"github.com/devlights/gomy/output"
)

// KeepAllocatedMemory -- スライスのメモリ状態をキープしたままで len を 0 にするサンプルです.
//
// REFERENCES:
//   - https://yourbasic.org/golang/clear-slice/
func KeepAllocatedMemory() error {
	// -------------------------------------------------------
	// スライスのメモリ状態をキープしたまま len を 0 にするには
	// 元のスライスの先頭を指し示すスライスを代入する.
	// つまり、意図的にサイズを0にするだけ.
	//
	// これにより、len は 0 となるが、cap は元のままとなる.
	// 当然であるが、範囲を広げると元のデータが見える.
	// -------------------------------------------------------
	s0 := []int{1, 2, 3}
	output.Stdoutf("[s0]", "%v\tlen=%d\tcap=%d\n", s0, len(s0), cap(s0))

	s1 := s0[:0]
	output.Stdoutf("[s1]", "%v\tlen=%d\tcap=%d\n", s1, len(s1), cap(s1))

	// 結果は
	// [s0]                 [1 2 3]	len=3	cap=3
	// [s1]                 []	len=0	cap=3
	// となる

	// 元のメモリ状態はそのまま残っているが、現在の s1 は、len=0 のスライスとなっている
	// なので、以下のようにインデックスの開始を範囲外にしてスライスを得ようとしても panic する
	/*
		s2 := s1[1:]
		output.Stdoutl("[s1[1:]]", s2)
	*/
	// panic: runtime error: slice bounds out of range [1:0]

	// スライス自体の範囲を広げることは可能
	// なので、以下のようにすると元のデータが見える.
	s3 := s1[:cap(s1)-1]
	output.Stdoutf("[s3]", "%v\tlen=%d\tcap=%d\n", s3, len(s3), cap(s3))

	// 結果は
	// [s3]                 [1 2]	len=2	cap=3
	// となる

	// では、ここから len=0 のスライスとして存在している s1 に
	// append するとどのようになるかを確認.
	//
	// s1 は、 元々のスライス s0 ([]int{1, 2, 3}) の先頭を指している.
	output.StdoutHr()
	output.Stdoutl("[append(s1, 4)]")

	s1 = append(s1, 4)
	output.Stdoutf("[result]", "s0=%v\ts1=%v\ts3=%v\n", s0, s1, s3)

	// 結果は、 s0=[4 2 3]	s1=[4]	s3=[4 2] となり
	// s1 には、新たに要素 4 が追加され、元のスライス s0 は 0番目 の要素が上書きで 4 に変わる
	// s3 も s0 と同じメモリ領域を指している別のスライスなので、同じように 0番目 が上書きで変わる

	// では、次に s3 に対して append するとどうなるかを確認.
	output.StdoutHr()
	output.Stdoutl("[append(s3, 5)]")

	s3 = append(s3, 5)
	output.Stdoutf("[result]", "s0=%v\ts1=%v\ts3=%v\n", s0, s1, s3)

	// 結果は、 s0=[4 2 5]	s1=[4]	s3=[4 2 5] となる
	// さらに s3 に append してみる
	output.StdoutHr()
	output.Stdoutl("[append(s3, 6)]")

	s3 = append(s3, 6)
	output.Stdoutf("[result]", "s0=%v\ts1=%v\ts3=%v\n", s0, s1, s3)

	// 結果は、s0=[4 2 5]	s1=[4]	s3=[4 2 5 6] となり、今度は s3 のみが変化する
	// これは、元のスライスの cap が 3 であり、今回の append によって元のスライスの cap を
	// 超えたため、別のスライスが割り当てられたため。なので、この段階で s3 は s0, s1 と別のスライス
	// となっている。

	// それを確認するために、 s1 にさらに append してみる
	// s1 の cap は、元のスライス s0 の状態であるので、 cap=3 である
	// 現在の s1 は、len=1 なので、まだ cap に到達していない
	output.StdoutHr()
	output.Stdoutl("[append(s1, 7)]")

	s1 = append(s1, 7)
	output.Stdoutf("[result]", "s0=%v\ts1=%v\ts3=%v\n", s0, s1, s3)

	// 結果は、s0=[4 7 5]	s1=[4 7]	s3=[4 2 5 6] となり、予想通り s3 は変化しない
	// 次に、元のスライス s0 に対して append を行い、capを超えて別のスライスになることを確認する
	output.StdoutHr()
	output.Stdoutl("[append(s0, 8)]")

	s0 = append(s0, 8)
	output.Stdoutf("[result]", "s0=%v\ts1=%v\ts3=%v\n", s0, s1, s3)

	// 結果は、s0=[4 7 5 8]	s1=[4 7]	s3=[4 2 5 6] となる.
	// スライス s0 は、今回の append で cap を超えたので、新たなスライスが割り当てられている
	// なので、 s0 と s1 は別のスライスとなっている.

	// ここまでで、 s0, s1, s3 は、元々は一つのスライスのメモリ領域を見ていた状態から
	// 全部バラバラのメモリ領域を見るようになっているはずである。
	// なので、最後に ３つのスライスに対して append してみる。
	output.StdoutHr()
	output.Stdoutl("[append s0, s1, s3]")

	s0 = append(s0, 10)
	s1 = append(s1, 11)
	s3 = append(s3, 12)

	output.Stdoutf("[result]", "s0=%v\ts1=%v\ts3=%v\n", s0, s1, s3)

	// 結果は、s0=[4 7 5 8 10]	s1=[4 7 11]	s3=[4 2 5 6 12] となる.
	// 予想通り、全部別のメモリ領域を指しているため、3つのスライスに対しての append は
	// それぞれのスライスにしか反映されない.

	return nil
}
