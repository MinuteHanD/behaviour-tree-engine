package bt

import "testing"

func TestTreeWithoutRootFailsSafely(t *testing.T) {
	tree := NewTree(nil)
	if got := tree.Tick(&Context{}); got != Failure {
		t.Fatalf("Tick() = %s, want %s", got, Failure)
	}

	tree.Reset()
}

func TestStatusString(t *testing.T) {
	tests := map[Status]string{
		Success:   "success",
		Failure:   "failure",
		Running:   "running",
		Status(9): "unknown",
	}

	for status, want := range tests {
		if got := status.String(); got != want {
			t.Errorf("Status(%d).String() = %q, want %q", status, got, want)
		}
	}
}

func TestWalk(t *testing.T) {
	child := NodeFunc(func(*Context) Status { return Success })
	root := walkNode{child: child}
	seen := 0

	Walk(root, func(Node) bool {
		seen++
		return true
	})
	if seen != 2 {
		t.Fatalf("visited %d nodes, want 2", seen)
	}
}

type walkNode struct {
	child Node
}

func (n walkNode) Tick(*Context) Status {
	return n.child.Tick(nil)
}

func (walkNode) Reset() {}

func (n walkNode) Child() Node {
	return n.child
}
