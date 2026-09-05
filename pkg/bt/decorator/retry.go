package decorator

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Retry struct {
	child       bt.Node
	maxAttempts int
	attempts    int
}

func NewRetry(child bt.Node, maxAttempts int) *Retry {
	if maxAttempts < 0 {
		maxAttempts = 0
	}
	return &Retry{child: child, maxAttempts: maxAttempts}
}

func (r *Retry) Child() bt.Node {
	return r.child
}

func (r *Retry) Tick(ctx *bt.Context) bt.Status {
	if r.child == nil || r.maxAttempts == 0 {
		return bt.Failure
	}

	status := r.child.Tick(ctx)

	if status == bt.Running {
		return bt.Running
	}

	if status == bt.Success {
		r.attempts = 0
		r.child.Reset()
		return bt.Success
	}

	r.attempts++
	r.child.Reset()

	if r.attempts >= r.maxAttempts {
		r.attempts = 0
		return bt.Failure
	}

	return bt.Running
}

func (r *Retry) Reset() {
	r.attempts = 0
	if r.child != nil {
		r.child.Reset()
	}
}
