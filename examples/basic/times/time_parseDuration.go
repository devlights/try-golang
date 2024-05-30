package times

import (
	"time"

	"github.com/devlights/gomy/output"
)

// ParseDuration は、time.ParseDuration() のサンプルです.
//
// > ParseDuration parses a duration string.
// A duration string is a possibly signed sequence of decimal numbers,
// each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
//
// > ParseDuration は継続時間文字列を解析する。
// 継続時間文字列は、符号付きの10進数列の可能性があり、
// それぞれオプションの分数と "300ms"、"-1.5h"、"2h45m" のような単位の接尾辞があります。
// 有効な時間単位は、"ns"、"us"（または "µs"）、"ms"、"s"、"m"、"h "である。
//
// # REFERENCES
//   - https://pkg.go.dev/time@go1.22.3#ParseDuration
func ParseDuration() error {
	var (
		items = []string{
			"500ms",
			"5s",
			"5m",
			"5h",
			"1h2m3s444ms",
			"1h2m3s444ms555us666ns", // ナノ秒まで指定
			"1h2m3s4d",              // 不正な接尾辞
			"1h2m3h4s5h",            // 同じ時間単位のものは合計される
			"1h2m3h4s5h500ms500ms",  // 同じ時間単位のものは合計される
			"-1h2m3h4s5h",           // 先頭に - を付与すると負の値に出来る
		}
	)

	for _, item := range items {
		var (
			d   time.Duration
			err error
		)

		d, err = time.ParseDuration(item)
		if err != nil {
			output.Stdoutl("[error]", err)
			continue
		}

		output.Stdoutl("[Duration]", d)
	}

	return nil
}
