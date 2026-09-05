package decorator

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Repeater struct {
	child      bt.Node
	maxRepeats int
	count      int
}

func NewRepeater(child bt.Node, maxRepeats int) *Repeater {
	if maxRepeats < 0 {
		maxRepeats = 0
	}
	return &Repeater{child: child, maxRepeats: maxRepeats}
}

func (r *Repeater) Child() bt.Node {
	return r.child
}

func (r *Repeater) Tick(ctx *bt.Context) bt.Status {
	if r.maxRepeats == 0 {
		return bt.Success
	}
	if r.child == nil {
		return bt.Failure
	}

	status := r.child.Tick(ctx)

	if status == bt.Running {
		return bt.Running
	}

	r.count++
	r.child.Reset()

	if r.count >= r.maxRepeats {
		r.count = 0
		return bt.Success
	}

	return bt.Running
}

func (r *Repeater) Reset() {
	r.count = 0
	if r.child != nil {
		r.child.Reset()
	}
}
