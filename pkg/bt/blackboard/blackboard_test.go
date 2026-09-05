package blackboard

import "testing"

func TestZeroValueIsUsable(t *testing.T) {
	var board Blackboard
	board.Set("answer", 42)

	value, ok := board.Get("answer")
	if !ok || value != 42 {
		t.Fatalf("Get(answer) = (%v, %t), want (42, true)", value, ok)
	}
}

func TestDeleteClearSnapshotAndTypedGet(t *testing.T) {
	board := New()
	board.Set("count", 2)
	board.Set("name", "guard")

	count, ok := Get[int](board, "count")
	if !ok || count != 2 {
		t.Fatalf("Get[int](count) = (%d, %t), want (2, true)", count, ok)
	}
	if _, ok := Get[int](board, "name"); ok {
		t.Fatal("type mismatch should not succeed")
	}

	snapshot := board.Snapshot()
	snapshot["count"] = 9
	count, _ = Get[int](board, "count")
	if count != 2 {
		t.Fatalf("snapshot mutated board: count = %d", count)
	}

	board.Delete("name")
	if board.Len() != 1 {
		t.Fatalf("Len() = %d, want 1", board.Len())
	}
	board.Clear()
	if board.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", board.Len())
	}
}
