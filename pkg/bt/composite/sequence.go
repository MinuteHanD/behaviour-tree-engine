package composite

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Sequence struct {
	children []bt.Node
	current  int
}

func NewSequence(children ...bt.Node) *Sequence {
	return &Sequence{children: children}
}

func (s *Sequence) Children() []bt.Node {
	return append([]bt.Node(nil), s.children...)
}

func (s *Sequence) Tick(ctx *bt.Context) bt.Status {
	for s.current < len(s.children) {
		child := s.children[s.current]
		if child == nil {
			s.Reset()
			return bt.Failure
		}
		status := child.Tick(ctx)

		switch status {
		case bt.Success:
			s.current++
		case bt.Failure:
			s.Reset()
			return bt.Failure
		case bt.Running:
			return bt.Running
		}
	}

	s.Reset()
	return bt.Success
}

func (s *Sequence) Reset() {
	s.current = 0
	for _, child := range s.children {
		if child != nil {
			child.Reset()
		}
	}
}
