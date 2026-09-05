package tests

import (
	"testing"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/composite"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

func TestSelector_WaitThenSuccess(t *testing.T) {
	wait := leaf.NewWait(2)
	actionRan := false
	action := leaf.NewAction(func(ctx *bt.Context) bt.Status {
		actionRan = true
		return bt.Success
	})

	selector := composite.NewSelector(wait, action)
	ctx := &bt.Context{}

	status := selector.Tick(ctx)
	if status != bt.Running {
		t.Errorf("expected running got %v", status)
	}

	if actionRan {
		t.Errorf("action should not have ran")
	}

	status = selector.Tick(ctx)
	if status != bt.Running {
		t.Errorf("expected running got %v", status)
	}

	if actionRan {
		t.Errorf("action should not have ran")
	}

	status = selector.Tick(ctx)
	if status != bt.Success {
		t.Errorf("expected success got %v", status)
	}

	if actionRan {
		t.Errorf("action should not have ran")
	}
}
