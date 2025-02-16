package strs

import "github.com/devlights/try-golang/mapping"

type (
	register struct{}
)

// NewRegister -- このパッケージ用のサンプルを登録する mapping.Register を生成します。
func NewRegister() mapping.Register {
	return new(register)
}

// Regist -- 登録します.
func (r *register) Regist(m mapping.ExampleMapping) {
	m["string_rune_rawstring"] = RuneRawString
	m["string_to_runeslice"] = ToRuneSlice
	m["string_rune_byte_convert"] = RuneByteConvert
	m["string_chop_newline"] = ChopNewLine
	m["string_using_builder"] = UsingBuilder
	m["string_rune_count"] = RuneCount
	m["string_diff_trimright_trimsuffix"] = DiffTrimRightAndTrimSuffix
	m["string_cut_prefix_suffix"] = CutPrefixSuffix
	m["string_using_clone"] = UsingStringsClone
	m["string_trim_space"] = TrimSpace
	m["string_random_string"] = RandomString
	m["string_split_fields"] = SplitFields
	m["string_go124_strings_lines"] = Go124StringsLines
}
