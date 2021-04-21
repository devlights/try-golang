package types

// MemoryPadding は、メンバー定義順によってメモリのパディングが発生する構造体です.
type MemoryPadding struct {
	Flag1    bool
	ShortVal int16
	Flag2    bool
	FloatVal float32
}

// NoMemoryPadding は、メンバー定義順を考慮してメモリのパティングが発生しないようにしている構造体です.
type NoMemoryPadding struct {
	FloatVal float32
	ShortVal int16
	Flag1    bool
	Flag2    bool
}

func (MemoryPadding) Layout() string {
	return `
	| Flag1    |             | ShortVal  | Flag2    |             | FloatVal    |
	-----------------------------------------------------------------------------
	| bool (1) | padding (1) | int16 (2) | bool (1) | padding (3) | float32 (4) |
	|                  4                 |          4             |      4      |
	`
}

func (NoMemoryPadding) Layout() string {
	return `
	| FloatVal    | ShortVal  | Flag1    | Flag2    |
	-------------------------------------------------
	| float32 (4) | int16 (2) | bool (1) | bool (1) |
	|     4       |              4                  |
	`
}
