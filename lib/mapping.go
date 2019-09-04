package lib

import (
	"github.com/devlights/try-golang/advanced/async"
	"github.com/devlights/try-golang/advanced/closure"
	"github.com/devlights/try-golang/advanced/reflection"
	"github.com/devlights/try-golang/advanced/sets"
	"github.com/devlights/try-golang/basic/array_"
	"github.com/devlights/try-golang/basic/comments"
	"github.com/devlights/try-golang/basic/constants"
	"github.com/devlights/try-golang/basic/defer_"
	"github.com/devlights/try-golang/basic/error_"
	"github.com/devlights/try-golang/basic/functions"
	"github.com/devlights/try-golang/basic/helloworld"
	"github.com/devlights/try-golang/basic/import_"
	"github.com/devlights/try-golang/basic/interface_"
	"github.com/devlights/try-golang/basic/io_"
	"github.com/devlights/try-golang/basic/iota_"
	"github.com/devlights/try-golang/basic/literals"
	"github.com/devlights/try-golang/basic/map_"
	"github.com/devlights/try-golang/basic/math_"
	"github.com/devlights/try-golang/basic/os_"
	"github.com/devlights/try-golang/basic/runtime_"
	"github.com/devlights/try-golang/basic/scope"
	"github.com/devlights/try-golang/basic/slice_"
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"github.com/devlights/try-golang/basic/string_"
	"github.com/devlights/try-golang/basic/struct_"
	"github.com/devlights/try-golang/basic/tutorial"
	"github.com/devlights/try-golang/basic/type_"
	"github.com/devlights/try-golang/basic/variables"
)

// SampleMappingは、サンプル名とサンプル呼び出し関数のマッピング定義を持つ型です
type SampleMapping map[string]func() error

// NewSampleMapping は、SampleMappingのコンストラクタ関数です
func NewSampleMapping() SampleMapping {
	return make(SampleMapping)
}

// MakeMapping は、マッピング生成します
func (m SampleMapping) MakeMapping() {
	m["error01"] = error_.Error01
	m["helloworld"] = helloworld.HelloWorld
	m["printf01"] = stdout.Printf01
	m["printf02"] = stdout.Printf02
	m["printf03"] = stdout.Printf03
	m["println01"] = stdout.Println01
	m["scanner01"] = stdin.Scanner01
	m["map_basic"] = map_.MapBasic
	m["map_for"] = map_.MapFor
	m["map_initialize"] = map_.MapInitialize
	m["map_delete"] = map_.MapDelete
	m["map_access"] = map_.MapAccess
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
	m["slice_reverse"] = slice_.SliceReverse
	m["comment01"] = comments.Comment01
	m["closure01"] = closure.Closure01
	m["string_rune_rawstring"] = string_.StringRuneRawString
	m["string_to_runeslice"] = string_.StringToRuneSlice
	m["set01"] = sets.Set01
	m["set02"] = sets.Set02
	m["set03"] = sets.Set03
	m["set04"] = sets.Set04
	m["set05"] = sets.Set05
	m["mapset01"] = sets.MapSet01
	m["defer01"] = defer_.Defer01
	m["var_statement_declare"] = variables.VarStatementDeclares
	m["package_scope_variable"] = variables.PackageScopeVariable
	m["short_assignment_statement"] = variables.ShortAssignmentStatement
	m["const_statement_declare"] = constants.ConstStatementDeclares
	m["function_one_return_value"] = functions.FunctionOneReturnValue
	m["function_multi_return_value"] = functions.FunctionMultiReturnValue
	m["function_named_return_value"] = functions.FunctionNamedReturnValue
	m["type01"] = type_.Type01
	m["minmax"] = math_.MinMax
	m["binary_int_literals"] = literals.BinaryIntLiterals
	m["octal_int_literals"] = literals.OctalIntLiterals
	m["hex_int_literals"] = literals.HexIntLiterals
	m["digit_separator"] = literals.DigitSeparators

	m["tutorial_gotour_helloworld"] = tutorial.GoTourHelloWorld
	m["tutorial_gotour_import"] = tutorial.GoTourImport
	m["tutorial_gotour_scope"] = tutorial.GoTourScope
	m["tutorial_gotour_functions"] = tutorial.GoTourFunctions
	m["tutorial_gotour_basictypes"] = tutorial.GoTourBasicTypes
	m["tutorial_gotour_zerovalue"] = tutorial.GoTourZeroValue
	m["tutorial_gotour_typeconvert_basictypes"] = tutorial.GoTourTypeConvertBasicTypes
}
