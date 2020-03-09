package enumerable

type (
	// Range は、範囲を表すインターフェースです。
	Range interface {
		// 開始点の値を返します.
		Start() int
		// 終了点の値を返します.
		End() int
		// 次の値に進みます. 進むことが出来ない場合は false を返します.
		Next() bool
		// 現在の値を返します.
		Current() int
		// 現在の値をリセットして開始点の値に戻します. 戻り値は (リセット直前の値, 処理で発生したエラー) です.
		Reset() (int, error)
	}

	enumerableRange struct {
		start, end, current int
	}
)

// NewRange は、指定された値を元に Range を生成して返します.
func NewRange(start, end int) Range {
	return &enumerableRange{
		start:   start,
		end:     end,
		current: start,
	}
}

func (e *enumerableRange) Start() int {
	return e.start
}

func (e *enumerableRange) End() int {
	return e.end
}

func (e *enumerableRange) Next() bool {
	if e.current == e.end {
		return false
	}

	e.current++
	return true
}

func (e *enumerableRange) Current() int {
	if e.start == e.current {
		return e.current
	}

	return e.current - 1
}

func (e *enumerableRange) Reset() (int, error) {
	cur := e.current
	e.current = e.start
	return cur, nil
}
