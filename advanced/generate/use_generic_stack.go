package generate

import (
	"github.com/devlights/try-golang/lib/output"
)

func UseGenericStack() error {
	bStack := NewBoolStack()
	bStack.Push(false)
	bStack.Push(true)

	sStack := NewStringStack()
	sStack.Push("world")
	sStack.Push("hello")

	iStack := NewIntStack()
	iStack.Push(100)
	iStack.Push(99)

	output.Stdoutl("[BoolStack]", bStack.Pop(), bStack.Pop())
	output.Stdoutl("[StringStack]", sStack.Pop(), sStack.Pop())
	output.Stdoutl("[IntStack]", iStack.Pop(), iStack.Pop())

	return nil
}
