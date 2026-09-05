package decorator

import (
	"time"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Timeout struct {
	child    bt.Node
	duration time.Duration
	elapsed  time.Duration
}

func NewTimeout(child bt.Node, duration time.Duration) *Timeout {
	if duration < 0 {
		duration = 0
	}
	return &Timeout{child: child, duration: duration}
}

func (d *Timeout) Tick(ctx *bt.Context) bt.Status {
	if d.child == nil {
		return bt.Failure
	}
	if d.elapsed >= d.duration {
		d.Reset()
		return bt.Failure
	}

	status := d.child.Tick(ctx)
	if status != bt.Running {
		d.Reset()
		return status
	}
	if ctx != nil && ctx.DeltaTime > 0 {
		d.elapsed += ctx.DeltaTime
	}
	return bt.Running
}

func (d *Timeout) Reset() {
	d.elapsed = 0
	if d.child != nil {
		d.child.Reset()
	}
}

func (d *Timeout) Child() bt.Node {
	return d.child
}
