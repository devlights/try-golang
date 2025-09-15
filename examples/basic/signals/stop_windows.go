package signals

import (
	"fmt"
)

func Stop() error {
	fmt.Println("Linux版のみです。")
	return nil
}
