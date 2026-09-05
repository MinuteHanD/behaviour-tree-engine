package composite

import (
	"testing"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

func TestReactiveSelectorRechecksHigherPriorityChild(t *testing.T) {
	priority := false
	fallbackCalls := 0
	selector := NewReactiveSelector(
		leaf.NewCondition(func(*bt.Context) bool { return priority }),
		leaf.NewAction(func(*bt.Context) bt.Status {
			fallbackCalls++
			return bt.Running
		}),
	)

	if got := selector.Tick(nil); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	priority = true
	if got := selector.Tick(nil); got != bt.Success {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Success)
	}
	if fallbackCalls != 1 {
		t.Fatalf("fallback calls = %d, want 1", fallbackCalls)
	}
}

func TestReactiveSequenceResetsLaterChildWhenGuardStopsPassing(t *testing.T) {
	allowed := true
	resets := 0
	child := &resetNode{status: bt.Running, resets: &resets}
	sequence := NewReactiveSequence(
		leaf.NewCondition(func(*bt.Context) bool { return allowed }),
		child,
	)

	if got := sequence.Tick(nil); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	allowed = false
	if got := sequence.Tick(nil); got != bt.Failure {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Failure)
	}
	if resets == 0 {
		t.Fatal("running child was not reset")
	}
}

type resetNode struct {
	status bt.Status
	resets *int
}

func (n *resetNode) Tick(*bt.Context) bt.Status {
	return n.status
}

func (n *resetNode) Reset() {
	*n.resets++
}
