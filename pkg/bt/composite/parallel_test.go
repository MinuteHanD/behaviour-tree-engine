package composite

import (
	"testing"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

func TestParallelRetainsCompletedChildren(t *testing.T) {
	ticks := 0
	runningThenSuccess := leaf.NewAction(func(*bt.Context) bt.Status {
		ticks++
		if ticks == 1 {
			return bt.Running
		}
		return bt.Success
	})
	alreadySuccessful := leaf.NewAction(func(*bt.Context) bt.Status { return bt.Success })

	parallel := NewParallel(2, alreadySuccessful, runningThenSuccess)
	if got := parallel.Tick(&bt.Context{}); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	if got := parallel.Tick(&bt.Context{}); got != bt.Success {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Success)
	}
}

func TestParallelImpossibleThresholdFails(t *testing.T) {
	parallel := NewParallel(2, leaf.NewAction(func(*bt.Context) bt.Status { return bt.Success }))
	if got := parallel.Tick(&bt.Context{}); got != bt.Failure {
		t.Fatalf("Tick() = %s, want %s", got, bt.Failure)
	}
}
