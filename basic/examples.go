package basic

import (
	"github.com/devlights/try-golang/advanced/sets"
	"github.com/devlights/try-golang/basic/array_"
	"github.com/devlights/try-golang/basic/builtin_"
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
	"github.com/devlights/try-golang/basic/log_"
	"github.com/devlights/try-golang/basic/map_"
	"github.com/devlights/try-golang/basic/math_"
	"github.com/devlights/try-golang/basic/os_"
	"github.com/devlights/try-golang/basic/runtime_"
	"github.com/devlights/try-golang/basic/scope"
	"github.com/devlights/try-golang/basic/slice_"
	"github.com/devlights/try-golang/basic/sort_"
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"github.com/devlights/try-golang/basic/strconv_"
	"github.com/devlights/try-golang/basic/string_"
	"github.com/devlights/try-golang/basic/struct_"
	"github.com/devlights/try-golang/basic/time_"
	"github.com/devlights/try-golang/basic/type_"
	"github.com/devlights/try-golang/basic/unsafe_"
	"github.com/devlights/try-golang/basic/variables"
	"github.com/devlights/try-golang/interfaces"
)

type (
	basicExampleRegister struct{}
)

// NewRegister は、basic パッケージ用の lib.Register を返します.
func NewRegister() interfaces.Register {
	r := new(basicExampleRegister)
	return r
}

// Regist は、basic パッケージ配下に存在するサンプルを登録します.
func (r *basicExampleRegister) Regist(m interfaces.ExampleMapping) {
	m["builtin_print"] = builtin_.PrintFunc
	m["error_basic"] = error_.Basic
	m["error_sentinel"] = error_.Sentinel
	m["error_typeassertion"] = error_.TypeAssertion
	m["error_wrap_unwrap"] = error_.WrapAndUnwrap
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
	m["import01"] = import_.Import01
	m["iota01"] = iota_.Iota01
	m["fileio01"] = io_.FileIo01
	m["fileio02"] = io_.FileIo02
	m["fileio03"] = io_.FileIo03
	m["fileio04"] = io_.FileIo04
	m["interface_basic"] = interface_.Basic
	m["interface_composition"] = interface_.Composition
	m["interface_ducktyping"] = interface_.DuckTyping
	m["os01"] = os_.Os01
	m["runtime_version"] = runtime_.RuntimeVersion
	m["runtime_memorystats"] = runtime_.RuntimeMemoryStats
	m["struct01"] = struct_.Struct01
	m["struct02"] = struct_.Struct02
	m["struct03"] = struct_.Struct03
	m["struct04"] = struct_.Struct04
	m["struct_anonymous_struct"] = struct_.StructAnonymousStruct
	m["struct_empty_struct"] = struct_.EmptyStruct
	m["array01"] = array_.Array01
	m["slice01"] = slice_.Slice01
	m["slice02"] = slice_.Slice02
	m["slice03"] = slice_.Slice03
	m["slice04"] = slice_.Slice04
	m["slice05"] = slice_.Slice05
	m["slice_reverse"] = slice_.SliceReverse
	m["slice_append"] = slice_.SliceAppend
	m["slice_pointer"] = slice_.SlicePointer
	m["slice_copy"] = slice_.SliceCopy
	m["slice_clear"] = slice_.SliceClear
	m["comment01"] = comments.Comment01
	m["string_rune_rawstring"] = string_.StringRuneRawString
	m["string_to_runeslice"] = string_.StringToRuneSlice
	m["string_rune_byte_convert"] = string_.StringRuneByteConvert
	m["mapset01"] = sets.MapSet01
	m["defer01"] = defer_.Defer01
	m["defer_in_loop"] = defer_.DeferInLoop
	m["var_statement_declare"] = variables.VarStatementDeclares
	m["package_scope_variable"] = variables.PackageScopeVariable
	m["short_assignment_statement"] = variables.ShortAssignmentStatement
	m["shadowing_variable"] = variables.ShadowingVariable
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
	m["time_since"] = time_.TimeSince
	m["time_after"] = time_.TimeAfter
	m["time_unix_to_time"] = time_.TimeUnixToTime
	m["time_now"] = time_.TimeNow
	m["time_parse"] = time_.TimeParse
	m["hex_to_decimal_convert"] = strconv_.HexToDecimalConvert
	m["unsafe_sizeof"] = unsafe_.Sizeof
	m["log_flags"] = log_.Flags
	m["log_prefix"] = log_.Prefix
	m["log_sentry_basic"] = log_.SentryBasic
	m["log_sentry_goroutine_bad"] = log_.SentryGoroutineBad
	m["log_sentry_goroutine_good"] = log_.SentryGoroutineGood
	m["log_output"] = log_.Output
	m["log_new"] = log_.NewLogger
	m["sort_interface"] = sort_.SortInterface
}
