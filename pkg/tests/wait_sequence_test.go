package tests

import (
	"testing"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/composite"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

func TestSequence_WaitThenAction(t *testing.T) {
	wait := leaf.NewWait(2)

	actionRan := false
	action := leaf.NewAction(func(ctx *bt.Context) bt.Status {
		actionRan = true
		return bt.Success
	})

	seq := composite.NewSequence(wait, action)
	ctx := &bt.Context{}

	status := seq.Tick(ctx)
	if status != bt.Running {
		t.Errorf("Expected Running, got %v", status)
	}

	if actionRan {
		t.Errorf("action should not have ran yet")
	}

	status = seq.Tick(ctx)
	if status != bt.Running {
		t.Errorf("Expected Running, got %v", status)
	}

	if actionRan {
		t.Errorf("action should not have ran yet")
	}

	status = seq.Tick(ctx)
	if status != bt.Success {
		t.Errorf("Expected Success, got %v", status)
	}

	if !actionRan {
		t.Errorf("action should have run")
	}

}
