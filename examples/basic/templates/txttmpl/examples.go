package txttmpl

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
	m["templates_text_tmpl_new"] = New
	m["templates_text_tmpl_must"] = Must
	m["templates_text_tmpl_structure"] = Structure
	m["templates_text_tmpl_whitespace_and_minus"] = WhitespaceAndMinussign
	m["templates_text_tmpl_comment"] = Comment
	m["templates_text_tmpl_if"] = If
	m["templates_text_tmpl_elseif"] = ElseIf
	m["templates_text_tmpl_else"] = Else
	m["templates_text_tmpl_range"] = Range
	m["templates_text_tmpl_range_else"] = RangeElse
	m["templates_text_tmpl_newline"] = Newline
	m["templates_text_tmpl_eq"] = Eq
	m["templates_text_tmpl_ne"] = Ne
	m["templates_text_tmpl_lt"] = Lt
}
