package leaf

import "github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"

type Named struct {
	Name  string
	child bt.Node
}

func NewNamed(name string, child bt.Node) *Named {
	return &Named{Name: name, child: child}
}

func (n *Named) Tick(ctx *bt.Context) bt.Status {
	if n.child == nil {
		return bt.Failure
	}
	return n.child.Tick(ctx)
}

func (n *Named) Reset() {
	if n.child != nil {
		n.child.Reset()
	}
}

func (n *Named) Child() bt.Node {
	return n.child
}
