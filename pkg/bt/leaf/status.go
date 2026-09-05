package leaf

import "github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"

type StatusNode struct {
	status bt.Status
}

func NewStatus(status bt.Status) *StatusNode {
	return &StatusNode{status: status}
}

func Success() *StatusNode {
	return NewStatus(bt.Success)
}

func Failure() *StatusNode {
	return NewStatus(bt.Failure)
}

func Running() *StatusNode {
	return NewStatus(bt.Running)
}

func (n *StatusNode) Tick(*bt.Context) bt.Status {
	return n.status
}

func (*StatusNode) Reset() {}
