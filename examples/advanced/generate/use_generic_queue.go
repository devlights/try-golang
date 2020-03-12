package generate

import (
	"github.com/devlights/gomy/output"
	"github.com/devlights/try-golang/examples/advanced/generate/queue"
)

func UseGenericQueue() error {

	bQueue := queue.NewBoolQueue()
	output.Stdoutl("[BoolQueue][1]", bQueue.Enqueue(true))
	output.Stdoutl("[BoolQueue][2]", bQueue.Enqueue(false))
	output.Stdoutl("[BoolQueue][3]", bQueue.Enqueue(true))
	output.Stdoutl("[BoolQueue][Count]", bQueue.Count())

	for {
		v, ok := bQueue.Dequeue()
		output.Stdoutl("[BoolQueue][Value]", v, ok)
		if !ok {
			break
		}
	}

	sQueue := queue.NewStringQueue()
	output.Stdoutl("[StrQueue][1]", sQueue.Enqueue("hello"))
	output.Stdoutl("[StrQueue][2]", sQueue.Enqueue("world"))
	output.Stdoutl("[StrQueue][Count]", sQueue.Count())

	for {
		v, ok := sQueue.Dequeue()
		output.Stdoutl("[StrQueue][Value]", v, ok)
		if !ok {
			break
		}
	}

	return nil
}
