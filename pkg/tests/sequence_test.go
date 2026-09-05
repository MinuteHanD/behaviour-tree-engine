package tests

import (
	"testing"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/composite"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

func TestSequence_AllSuccess(t *testing.T) {
	action1 := leaf.NewAction(func(ctx *bt.Context) bt.Status { return bt.Success })
	action2 := leaf.NewAction(func(ctx *bt.Context) bt.Status { return bt.Success })

	seq := composite.NewSequence(action1, action2)

	ctx := &bt.Context{}

	status := seq.Tick(ctx)

	if status != bt.Success {
		t.Errorf("Expected Success, got %v", status)
	}
}

func TestSequence_FailsOnFirstFailure(t *testing.T) {
	failAction := leaf.NewAction(func(ctx *bt.Context) bt.Status { return bt.Failure })

	neverRunsAction := leaf.NewAction(func(ctx *bt.Context) bt.Status {
		t.Errorf("this should not run since the first actions failed")
		return bt.Success
	})

	seq := composite.NewSequence(failAction, neverRunsAction)
	ctx := &bt.Context{}

	status := seq.Tick(ctx)

	if status != bt.Failure {
		t.Errorf("Expected Failure got %v", status)
	}
}

func TestSequence_Running(t *testing.T) {
	successAction := leaf.NewAction(func(ctx *bt.Context) bt.Status { return bt.Success })
	runningAction := leaf.NewAction(func(ctx *bt.Context) bt.Status { return bt.Running })

	seq := composite.NewSequence(successAction, runningAction)
	ctx := &bt.Context{}

	status := seq.Tick(ctx)

	if status != bt.Running {
		t.Errorf("Expected Running, got %v", status)
	}
}
