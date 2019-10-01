package tutorial

// 利用するパッケージのimportを記載する。
// 複数のimportがある場合は、以下のようにグループ化して記述することができる。
import (
	"fmt"
	"math"
	math2 "math" // alias設定
)

func GoTourImport() error {

	// ------------------------------------------------------------
	// import したパッケージの利用
	//   import したパッケージは、その名前で利用することが出来る。
	//   (python などと同様)
	//   alias 定義した名前も利用できる。
	// ------------------------------------------------------------
	fmt.Println(math.Pi)
	fmt.Println(math2.Pi)

	return nil
}
