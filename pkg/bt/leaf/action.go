package leaf

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type ActionFunc func(ctx *bt.Context) bt.Status

type Action struct {
	doWork ActionFunc
}

func NewAction(fn ActionFunc) *Action {
	return &Action{doWork: fn}
}

func (a *Action) Tick(ctx *bt.Context) bt.Status {
	if a.doWork == nil {
		return bt.Failure
	}
	return a.doWork(ctx)
}

func (a *Action) Reset() {
}
