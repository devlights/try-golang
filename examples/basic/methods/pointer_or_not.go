package methods

import "github.com/devlights/gomy/output"

type (
	_nonPointerReceiver struct {
		val int
	}

	_pointerReceiver struct {
		val int
	}
)

func (me _nonPointerReceiver) update(val int) {
	me.val = val //lint:ignore SA4005 サンプルなので意図的にこのようにしている
}

func (me *_pointerReceiver) update(val int) {
	me.val = val
}

// PointerOrNot は、メソッドのレシーバーをポインタで宣言するかしないかの違いについてのサンプルです。
func PointerOrNot() error {
	var (
		nonPointer = _nonPointerReceiver{}
		pointer    = _pointerReceiver{}
	)

	// これは内部の値が更新されない
	nonPointer.update(100)
	// これは内部の値が更新される
	pointer.update(100)

	output.Stdoutl("[non-pointer]", nonPointer.val)
	output.Stdoutl("[pointer    ]", pointer.val)

	return nil

	/*
	   $ task
	   task: [build] go build .
	   task: [run] ./try-golang -onetime

	   ENTER EXAMPLE NAME: methods_pointer_or_not

	   [Name] "methods_pointer_or_not"
	   [non-pointer]        0
	   [pointer    ]        100


	   [Elapsed] 35.15µs
	*/

}
