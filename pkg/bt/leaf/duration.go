package leaf

import (
	"time"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type WaitDuration struct {
	duration time.Duration
	elapsed  time.Duration
}

func NewWaitDuration(duration time.Duration) *WaitDuration {
	if duration < 0 {
		duration = 0
	}
	return &WaitDuration{duration: duration}
}

func (w *WaitDuration) Tick(ctx *bt.Context) bt.Status {
	if w.elapsed >= w.duration {
		return bt.Success
	}
	if ctx != nil && ctx.DeltaTime > 0 {
		w.elapsed += ctx.DeltaTime
	}
	return bt.Running
}

func (w *WaitDuration) Reset() {
	w.elapsed = 0
}
