package main

import (
	"fmt"
	"strings"
)

type (
	// BitVector は、ビットベクタを表す構造体です.
	BitVector struct {
		bits []uint32
		size int
	}
	opcode int
)

const (
	AND opcode = iota
	OR  opcode = iota
	XOR opcode = iota
)

// NewBitVector は、指定されたサイズのビットベクタを作成します.
func NewBitVector(size int) *BitVector {
	return &BitVector{
		bits: make([]uint32, (size+31)/32),
		size: size,
	}
}

// Set は、指定されたインデックス位置のビットを設定します.
func (me *BitVector) Set(index int, value bool) error {
	if index < 0 || index >= me.size {
		return fmt.Errorf("index out of range: %d", index)
	}

	var (
		aryIndex = index / 32
		bitIndex = uint(index % 32)
	)
	if value {
		me.bits[aryIndex] |= 1 << bitIndex // ビットオン (OR)
	} else {
		me.bits[aryIndex] &^= 1 << bitIndex // ビットクリア (AND NOT)
	}

	return nil
}

// Get は、指定されたインデックス位置のビットを取得します.
func (me *BitVector) Get(index int) (bool, error) {
	if index < 0 || index >= me.size {
		return false, fmt.Errorf("index out of range: %d", index)
	}

	var (
		aryIndex = index / 32
		bitIndex = uint(index % 32)
		bit      = (me.bits[aryIndex] & (1 << bitIndex)) != 0
	)

	return bit, nil
}

// And は、指定された *BitVector とのAND結果を設定した新たな *BitVector を返します.
func (me *BitVector) And(other *BitVector) (*BitVector, error) {
	return me.calc(other, AND)
}

// Or は、指定された *BitVector とのOR結果を設定した新たな *BitVector を返します.
func (me *BitVector) Or(other *BitVector) (*BitVector, error) {
	return me.calc(other, OR)
}

// Xor は、指定された *BitVector とのXOR結果を設定した新たな *BitVector を返します.
func (me *BitVector) Xor(other *BitVector) (*BitVector, error) {
	return me.calc(other, XOR)
}

func (me *BitVector) calc(other *BitVector, op opcode) (*BitVector, error) {
	if me.size != other.size {
		return nil, fmt.Errorf("size mismatch: %d != %d", me.size, other.size)
	}

	var (
		result  = NewBitVector(me.size)
		arySize = (me.size + 31) / 32
	)
	for i := range arySize {
		switch op {
		case AND:
			result.bits[i] = me.bits[i] & other.bits[i]
		case OR:
			result.bits[i] = me.bits[i] | other.bits[i]
		case XOR:
			result.bits[i] = me.bits[i] ^ other.bits[i]
		default:
			return nil, fmt.Errorf("unknown kind: %d", op)
		}
	}

	return result, nil
}

// String は、ビットベクタの文字列表現を返します.
func (me *BitVector) String() string {
	var (
		sb  strings.Builder
		val bool
	)
	for i := range me.size {
		val, _ = me.Get(i)
		if val {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}

	return sb.String()
}
