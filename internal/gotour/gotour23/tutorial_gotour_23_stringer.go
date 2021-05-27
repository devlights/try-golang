package gotour23

import (
	"fmt"
)

type (
	// IPAddr -- IPアドレスを表します.
	IPAddr [4]byte
)

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

// 別回答
// func (i IPAddr) String() string {
//	var (
//		strs = make([]string, 0, len(i))
//	)
//
//	for _, v := range i {
//		strs = append(strs, fmt.Sprintf("%v", v))
//	}
//
//	return strings.Join(strs, ".")
// }

// Stringer は、 Tour of Go - Stringers (https://tour.golang.org/methods/17) の サンプルです。
func Stringer() error {
	// ------------------------------------------------------------
	// fmt.Stringer インターフェースは、最もよく利用されるインターフェース
	// の一つ。以下、A Tour of Go (http://bit.ly/34wWjIq) のエクササイズ
	// ------------------------------------------------------------
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	return nil
}
