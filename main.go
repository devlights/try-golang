package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/devlights/try-golang/advanced/async"
	"github.com/devlights/try-golang/advanced/closure"
	"github.com/devlights/try-golang/advanced/reflection"
	"github.com/devlights/try-golang/advanced/sets"
	"github.com/devlights/try-golang/basic/array_"
	"github.com/devlights/try-golang/basic/comments"
	"github.com/devlights/try-golang/basic/defer_"
	"github.com/devlights/try-golang/basic/error_"
	"github.com/devlights/try-golang/basic/helloworld"
	"github.com/devlights/try-golang/basic/import_"
	"github.com/devlights/try-golang/basic/interface_"
	"github.com/devlights/try-golang/basic/io_"
	"github.com/devlights/try-golang/basic/iota_"
	"github.com/devlights/try-golang/basic/map_"
	"github.com/devlights/try-golang/basic/os_"
	"github.com/devlights/try-golang/basic/runtime_"
	"github.com/devlights/try-golang/basic/scope"
	"github.com/devlights/try-golang/basic/slice_"
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"github.com/devlights/try-golang/basic/string_"
	"github.com/devlights/try-golang/basic/struct_"
	"github.com/devlights/try-golang/basic/variables"
	"log"
	"os"
	"strings"
)

// サンプル名とサンプル呼び出し関数のマッピング定義の型
type SampleMapping map[string]func() error

// マッピング生成
func (m SampleMapping) MakeMapping() {
	m["error01"] = error_.Error01
	m["helloworld"] = helloworld.HelloWorld
	m["printf01"] = stdout.Printf01
	m["printf02"] = stdout.Printf02
	m["printf03"] = stdout.Printf03
	m["println01"] = stdout.Println01
	m["scanner01"] = stdin.Scanner01
	m["map01"] = map_.Map01
	m["scope01"] = scope.Scope01
	m["async01"] = async.Async01
	m["reflection01"] = reflection.Reflection01
	m["import01"] = import_.Import01
	m["iota01"] = iota_.Iota01
	m["fileio01"] = io_.FileIo01
	m["fileio02"] = io_.FileIo02
	m["fileio03"] = io_.FileIo03
	m["fileio04"] = io_.FileIo04
	m["interface01"] = interface_.Interface01
	m["os01"] = os_.Os01
	m["runtime01"] = runtime_.Runtime01
	m["struct01"] = struct_.Struct01
	m["struct02"] = struct_.Struct02
	m["struct03"] = struct_.Struct03
	m["struct04"] = struct_.Struct04
	m["array01"] = array_.Array01
	m["slice01"] = slice_.Slice01
	m["slice02"] = slice_.Slice02
	m["slice03"] = slice_.Slice03
	m["slice04"] = slice_.Slice04
	m["slice05"] = slice_.Slice05
	m["comment01"] = comments.Comment01
	m["closure01"] = closure.Closure01
	m["string01"] = string_.String01
	m["set01"] = sets.Set01
	m["set02"] = sets.Set02
	m["set03"] = sets.Set03
	m["set04"] = sets.Set04
	m["set05"] = sets.Set05
	m["mapset01"] = sets.MapSet01
	m["defer01"] = defer_.Defer01
	m["var_statement_declare"] = variables.VarStatementDeclares
	m["package_scope_variable"] = variables.PackageScopeVariable
}

// サンプル関数のマッピング
var mapping = make(SampleMapping)

// 初期化関数
func init() {
	mapping.MakeMapping()
}

// メインエントリポイント
func main() {
	var (
		onetime = flag.Bool("onetime", false, "run only one time")
	)

	flag.Parse()

	fmt.Print("ENTER EXAMPLE NAME: ")

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		// 実行サンプル名取得
		example := stdinScanner.Text()
		if strings.ToLower(example) == "quit" {
			// 終了
			break
		}

		// サンプル実行
		if v, ok := mapping[example]; ok {
			if err := v(); err != nil {
				log.Fatal(err)
			}

			fmt.Print("\n\n")
		}

		if *onetime {
			break
		}

		fmt.Print("ENTER EXAMPLE NAME: ")
	}

	fmt.Println("END")
}
