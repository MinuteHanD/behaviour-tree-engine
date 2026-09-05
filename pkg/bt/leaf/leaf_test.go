package leaf

import (
	"testing"
	"time"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

func TestWaitDuration(t *testing.T) {
	wait := NewWaitDuration(20 * time.Millisecond)
	ctx := &bt.Context{DeltaTime: 10 * time.Millisecond}

	if got := wait.Tick(ctx); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	if got := wait.Tick(ctx); got != bt.Running {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Running)
	}
	if got := wait.Tick(ctx); got != bt.Success {
		t.Fatalf("third Tick() = %s, want %s", got, bt.Success)
	}
}

func TestStatusNodes(t *testing.T) {
	for _, test := range []struct {
		node bt.Node
		want bt.Status
	}{
		{Success(), bt.Success},
		{Failure(), bt.Failure},
		{Running(), bt.Running},
	} {
		if got := test.node.Tick(nil); got != test.want {
			t.Errorf("Tick() = %s, want %s", got, test.want)
		}
	}
}
