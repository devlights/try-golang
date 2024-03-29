package gotour02

// 利用するパッケージのimportを記載する。
// 複数のimportがある場合は、以下のようにグループ化して記述することができる。
import (
	"fmt"
	"math"
	myos "os"
)

// Import は、 Tour of Go - Imports (https://tour.golang.org/basics/2) の サンプルです。
func Import() error {

	// ------------------------------------------------------------
	// import したパッケージの利用
	//   import したパッケージは、その名前で利用することが出来る。
	//   (python などと同様)
	//   alias 定義した名前も利用できる。
	// ------------------------------------------------------------
	fmt.Println(math.Pi)
	fmt.Println(myos.ModePerm)

	return nil
}
