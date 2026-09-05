package bt

type Node interface {
	Tick(ctx *Context) Status
	Reset()
}

type NodeFunc func(ctx *Context) Status

func (fn NodeFunc) Tick(ctx *Context) Status {
	if fn == nil {
		return Failure
	}
	return fn(ctx)
}

func (NodeFunc) Reset() {}

type ChildNode interface {
	Child() Node
}

type ChildrenNode interface {
	Children() []Node
}

func Walk(root Node, visit func(Node) bool) {
	if root == nil || visit == nil || !visit(root) {
		return
	}

	switch node := root.(type) {
	case ChildNode:
		Walk(node.Child(), visit)
	case ChildrenNode:
		for _, child := range node.Children() {
			Walk(child, visit)
		}
	}
}
