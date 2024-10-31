package main

type (
	Gate struct {
		latch *CountdownLatch
	}
)

func NewGate() *Gate {
	var (
		latch = NewCountdownLatch(1)
		gate  = Gate{latch}
	)

	return &gate
}

func (me *Gate) Await() {
	if me.latch.CurrentCount() < 1 {
		return
	}

	me.latch.Wait()
}

func (me *Gate) Open() {
	if me.latch.CurrentCount() < 1 {
		return
	}

	me.latch.Signal()
}
