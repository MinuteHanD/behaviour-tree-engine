package decorator

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Inverter struct {
	child bt.Node
}

func NewInverter(child bt.Node) *Inverter {
	return &Inverter{child: child}
}

func (i *Inverter) Child() bt.Node {
	return i.child
}

func (i *Inverter) Tick(ctx *bt.Context) bt.Status {
	if i.child == nil {
		return bt.Failure
	}
	status := i.child.Tick(ctx)
	switch status {
	case bt.Success:
		i.child.Reset()
		return bt.Failure
	case bt.Failure:
		i.child.Reset()
		return bt.Success
	default:
		return status
	}
}

func (i *Inverter) Reset() {
	if i.child != nil {
		i.child.Reset()
	}
}
