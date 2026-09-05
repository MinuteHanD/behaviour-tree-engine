package bt

type Tree struct {
	root Node
}

func NewTree(root Node) *Tree {
	return &Tree{root: root}
}

func (t *Tree) Tick(ctx *Context) Status {
	if t == nil || t.root == nil {
		return Failure
	}
	if ctx == nil {
		ctx = &Context{}
	}
	return t.root.Tick(ctx)
}

func (t *Tree) Reset() {
	if t == nil || t.root == nil {
		return
	}
	t.root.Reset()
}

func (t *Tree) Root() Node {
	if t == nil {
		return nil
	}
	return t.root
}

func (t *Tree) Walk(visit func(Node) bool) {
	if t == nil {
		return
	}
	Walk(t.root, visit)
}
