package composite

import (
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Parallel struct {
	children        []bt.Node
	successRequired int
	completed       []bool
	statuses        []bt.Status
}

func NewParallel(successRequired int, children ...bt.Node) *Parallel {
	if successRequired < 0 {
		successRequired = 0
	}
	return &Parallel{
		children:        children,
		successRequired: successRequired,
		completed:       make([]bool, len(children)),
		statuses:        make([]bt.Status, len(children)),
	}
}

func NewParallelAll(children ...bt.Node) *Parallel {
	return NewParallel(len(children), children...)
}

func NewParallelAny(children ...bt.Node) *Parallel {
	return NewParallel(1, children...)
}

func (p *Parallel) Children() []bt.Node {
	return append([]bt.Node(nil), p.children...)
}

func (p *Parallel) Tick(ctx *bt.Context) bt.Status {
	if p.successRequired == 0 {
		p.Reset()
		return bt.Success
	}
	if p.successRequired > len(p.children) {
		p.Reset()
		return bt.Failure
	}

	successCount := 0
	failureCount := 0

	for index, child := range p.children {
		if p.completed[index] {
			if p.statuses[index] == bt.Success {
				successCount++
			} else {
				failureCount++
			}
			continue
		}

		if child == nil {
			p.completed[index] = true
			p.statuses[index] = bt.Failure
			failureCount++
			continue
		}

		status := child.Tick(ctx)
		if status == bt.Running {
			continue
		}

		p.completed[index] = true
		p.statuses[index] = status
		if status == bt.Success {
			successCount++
		} else {
			failureCount++
		}
	}

	if successCount >= p.successRequired {
		p.Reset()
		return bt.Success
	}

	if len(p.children)-failureCount < p.successRequired {
		p.Reset()
		return bt.Failure
	}

	return bt.Running
}

func (p *Parallel) Reset() {
	for index, child := range p.children {
		p.completed[index] = false
		p.statuses[index] = bt.Running
		if child != nil {
			child.Reset()
		}
	}
}
