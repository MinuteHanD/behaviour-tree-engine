package composite

import "github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"

type ReactiveSequence struct {
	children []bt.Node
}

func NewReactiveSequence(children ...bt.Node) *ReactiveSequence {
	return &ReactiveSequence{children: children}
}

func (s *ReactiveSequence) Tick(ctx *bt.Context) bt.Status {
	for index, child := range s.children {
		if child == nil {
			s.Reset()
			return bt.Failure
		}

		status := child.Tick(ctx)
		switch status {
		case bt.Success:
			continue
		case bt.Failure:
			s.Reset()
			return bt.Failure
		default:
			s.resetAfter(index)
			return bt.Running
		}
	}

	s.Reset()
	return bt.Success
}

func (s *ReactiveSequence) Reset() {
	for _, child := range s.children {
		if child != nil {
			child.Reset()
		}
	}
}

func (s *ReactiveSequence) Children() []bt.Node {
	return append([]bt.Node(nil), s.children...)
}

func (s *ReactiveSequence) resetAfter(index int) {
	for _, child := range s.children[index+1:] {
		if child != nil {
			child.Reset()
		}
	}
}

type ReactiveSelector struct {
	children []bt.Node
}

func NewReactiveSelector(children ...bt.Node) *ReactiveSelector {
	return &ReactiveSelector{children: children}
}

func (s *ReactiveSelector) Tick(ctx *bt.Context) bt.Status {
	for index, child := range s.children {
		if child == nil {
			continue
		}

		status := child.Tick(ctx)
		switch status {
		case bt.Failure:
			continue
		case bt.Success:
			s.Reset()
			return bt.Success
		default:
			s.resetAfter(index)
			return bt.Running
		}
	}

	s.Reset()
	return bt.Failure
}

func (s *ReactiveSelector) Reset() {
	for _, child := range s.children {
		if child != nil {
			child.Reset()
		}
	}
}

func (s *ReactiveSelector) Children() []bt.Node {
	return append([]bt.Node(nil), s.children...)
}

func (s *ReactiveSelector) resetAfter(index int) {
	for _, child := range s.children[index+1:] {
		if child != nil {
			child.Reset()
		}
	}
}
