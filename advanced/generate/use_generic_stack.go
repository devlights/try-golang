package generate

import (
	"github.com/devlights/try-golang/advanced/generate/stack"
	"github.com/devlights/try-golang/output"
)

func UseGenericStack() error {
	bStack := stack.NewBoolStack()
	bStack.Push(false)
	bStack.Push(true)

	sStack := stack.NewStringStack()
	sStack.Push("world")
	sStack.Push("hello")

	iStack := stack.NewIntStack()
	iStack.Push(100)
	iStack.Push(99)

	output.Stdoutl("[BoolStack]", bStack.Pop(), bStack.Pop())
	output.Stdoutl("[StringStack]", sStack.Pop(), sStack.Pop())
	output.Stdoutl("[IntStack]", iStack.Pop(), iStack.Pop())

	return nil
}
