package lib

import (
	"github.com/devlights/try-golang/advanced/async"
	"github.com/devlights/try-golang/advanced/closure"
	"github.com/devlights/try-golang/advanced/reflection"
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
	"github.com/devlights/try-golang/basic/stdin"
	"github.com/devlights/try-golang/basic/stdout"
	"github.com/devlights/try-golang/basic/strconv_"
	"github.com/devlights/try-golang/basic/string_"
	"github.com/devlights/try-golang/basic/struct_"
	"github.com/devlights/try-golang/basic/time_"
	"github.com/devlights/try-golang/basic/type_"
	"github.com/devlights/try-golang/basic/unsafe_"
	"github.com/devlights/try-golang/basic/variables"
	"github.com/devlights/try-golang/effectivego"
	"github.com/devlights/try-golang/tutorial"
)

type (
	SampleKey     string                   // SampleKeyは、サンプル名を表すキーを表します
	SampleFunc    func() error             // SampleFuncは、実行するサンプル処理を表します
	SampleMapping map[SampleKey]SampleFunc // SampleMappingは、サンプルのマッピング定義を表します
)

// NewSampleMapping は、SampleMappingのコンストラクタ関数です
func NewSampleMapping() SampleMapping {
	return make(SampleMapping)
}

// MakeMapping は、マッピングを生成します
func (m SampleMapping) MakeMapping() {
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
	m["async01"] = async.Async01
	m["reflection01"] = reflection.Reflection01
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
	m["runtime01"] = runtime_.Runtime01
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
	m["comment01"] = comments.Comment01
	m["closure01"] = closure.Closure01
	m["string_rune_rawstring"] = string_.StringRuneRawString
	m["string_to_runeslice"] = string_.StringToRuneSlice
	m["string_rune_byte_convert"] = string_.StringRuneByteConvert
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

	m["tutorial_gotour_helloworld"] = tutorial.HelloWorld
	m["tutorial_gotour_import"] = tutorial.Import
	m["tutorial_gotour_scope"] = tutorial.Scope
	m["tutorial_gotour_functions"] = tutorial.Functions
	m["tutorial_gotour_basictypes"] = tutorial.BasicTypes
	m["tutorial_gotour_zerovalue"] = tutorial.ZeroValue
	m["tutorial_gotour_typeconvert_basictypes"] = tutorial.TypeConvertBasicTypes
	m["tutorial_gotour_const"] = tutorial.Constant
	m["tutorial_gotour_forloop"] = tutorial.ForLoop
	m["tutorial_gotour_if"] = tutorial.If
	m["tutorial_gotour_switch"] = tutorial.Switch
	m["tutorial_gotour_defer"] = tutorial.Defer
	m["tutorial_gotour_pointer"] = tutorial.Pointer
	m["tutorial_gotour_struct"] = tutorial.Struct
	m["tutorial_gotour_array"] = tutorial.Array
	m["tutorial_gotour_slice"] = tutorial.Slice
	m["tutorial_gotour_map"] = tutorial.Map
	m["tutorial_gotour_method"] = tutorial.Method
	m["tutorial_gotour_interface"] = tutorial.Interface
	m["tutorial_gotour_empty_interface"] = tutorial.EmptyInterface
	m["tutorial_gotour_type_assertion"] = tutorial.TypeAssertion
	m["tutorial_gotour_type_switch"] = tutorial.TypeSwitch
	m["tutorial_gotour_stringer"] = tutorial.Stringer
	m["tutorial_gotour_error"] = tutorial.Error
	m["tutorial_gotour_reader"] = tutorial.Reader
	m["tutorial_gotour_goroutine"] = tutorial.Goroutine
	m["tutorial_gotour_channels"] = tutorial.Channels
	m["tutorial_gotour_select"] = tutorial.Select
	m["tutorial_gotour_mutex"] = tutorial.Mutex

	m["effective_go_intro"] = effectivego.Introduction
	m["effective_go_formatting"] = effectivego.Formatting
	m["effective_go_comment"] = effectivego.Commentary
	m["effective_go_names"] = effectivego.Names
	m["effective_go_semicolon"] = effectivego.Semicolons
	m["effective_go_control"] = effectivego.ControlStructure
	m["effective_go_functions"] = effectivego.Functions
	m["effective_go_defer"] = effectivego.Defer
	m["effective_go_allocation_with_new"] = effectivego.AllocationWithNew
	m["effective_go_constructors"] = effectivego.Constructors
	m["effective_go_allocation_with_make"] = effectivego.AllocationWithMake
	m["effective_go_arrays"] = effectivego.Arrays
	m["effective_go_slices"] = effectivego.Slices
	m["effective_go_two_dimentional_slices"] = effectivego.TwoDimentionalSlices
	m["effective_go_maps"] = effectivego.Maps
	m["effective_go_printing"] = effectivego.Printing
	m["effective_go_append"] = effectivego.Append
	m["effective_go_constants"] = effectivego.Constants
	m["effective_go_methods"] = effectivego.Methods
}
