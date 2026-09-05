package decorator

import (
	"testing"
	"time"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

func TestRepeaterRunsChildRequestedNumberOfTimes(t *testing.T) {
	calls := 0
	repeater := NewRepeater(leaf.NewAction(func(*bt.Context) bt.Status {
		calls++
		return bt.Failure
	}), 2)

	if got := repeater.Tick(&bt.Context{}); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	if got := repeater.Tick(&bt.Context{}); got != bt.Success {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Success)
	}
	if calls != 2 {
		t.Fatalf("child calls = %d, want 2", calls)
	}
}

func TestUntilSuccessRetriesFailures(t *testing.T) {
	calls := 0
	node := NewUntilSuccess(leaf.NewAction(func(*bt.Context) bt.Status {
		calls++
		if calls == 2 {
			return bt.Success
		}
		return bt.Failure
	}))

	if got := node.Tick(nil); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	if got := node.Tick(nil); got != bt.Success {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Success)
	}
}

func TestTimeoutFailsAfterDuration(t *testing.T) {
	node := NewTimeout(leaf.Running(), 10*time.Millisecond)
	ctx := &bt.Context{DeltaTime: 10 * time.Millisecond}

	if got := node.Tick(ctx); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	if got := node.Tick(ctx); got != bt.Failure {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Failure)
	}
}

func TestRetryStopsAfterMaximumAttempts(t *testing.T) {
	calls := 0
	retry := NewRetry(leaf.NewAction(func(*bt.Context) bt.Status {
		calls++
		return bt.Failure
	}), 2)

	if got := retry.Tick(&bt.Context{}); got != bt.Running {
		t.Fatalf("first Tick() = %s, want %s", got, bt.Running)
	}
	if got := retry.Tick(&bt.Context{}); got != bt.Failure {
		t.Fatalf("second Tick() = %s, want %s", got, bt.Failure)
	}
	if calls != 2 {
		t.Fatalf("child calls = %d, want 2", calls)
	}
}
