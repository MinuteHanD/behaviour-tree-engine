package leaf

import "github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"

type Wait struct {
	ticks          int
	ticksRemaining int
}

func NewWait(ticks int) *Wait {
	if ticks < 0 {
		ticks = 0
	}
	return &Wait{ticks: ticks, ticksRemaining: ticks}
}

func (w *Wait) Tick(_ *bt.Context) bt.Status {
	if w.ticksRemaining <= 0 {
		return bt.Success
	}

	w.ticksRemaining--
	return bt.Running
}

func (w *Wait) Reset() {
	w.ticksRemaining = w.ticks
}
