package composite

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Selector struct {
	children []bt.Node
	current  int
}

func NewSelector(children ...bt.Node) *Selector {
	return &Selector{children: children}
}

func (s *Selector) Children() []bt.Node {
	return append([]bt.Node(nil), s.children...)
}

func (s *Selector) Tick(ctx *bt.Context) bt.Status {
	for s.current < len(s.children) {
		child := s.children[s.current]
		if child == nil {
			s.current++
			continue
		}
		status := child.Tick(ctx)

		switch status {
		case bt.Success:
			s.Reset()
			return bt.Success
		case bt.Failure:
			s.current++
		case bt.Running:
			return bt.Running
		}
	}

	s.Reset()
	return bt.Failure
}

func (s *Selector) Reset() {
	s.current = 0
	for _, child := range s.children {
		if child != nil {
			child.Reset()
		}
	}
}
