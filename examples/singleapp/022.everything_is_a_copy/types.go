package main

import "fmt"

type ValueType struct {
	Value int
}

type PointerType struct {
	Value int
}

func NewValueType(v int) ValueType {
	o := ValueType{Value: v}
	return o
}

func NewPointerType(v int) *PointerType {
	o := &PointerType{Value: v}
	return o
}

func (me ValueType) Change() {
	me.Value += 100 //lint:ignore SA4005 It's ok because this is just a example.
}

func (me *PointerType) Change() {
	me.Value += 100
}

func (me ValueType) String() string {
	return fmt.Sprint(me.Value)
}

func (me *PointerType) String() string {
	return fmt.Sprint(me.Value)
}
