package leaf

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type ConditionFunc func(ctx *bt.Context) bool

type Condition struct {
	check ConditionFunc
}

func NewCondition(fn ConditionFunc) *Condition {
	return &Condition{check: fn}
}

func (c *Condition) Tick(ctx *bt.Context) bt.Status {
	if c.check == nil {
		return bt.Failure
	}
	if c.check(ctx) {
		return bt.Success
	}
	return bt.Failure
}

func (c *Condition) Reset() {
}
