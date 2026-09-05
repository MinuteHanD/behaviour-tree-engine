package decorator

import "github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"

type Succeeder struct {
	child bt.Node
}

func NewSucceeder(child bt.Node) *Succeeder {
	return &Succeeder{child: child}
}

func (d *Succeeder) Tick(ctx *bt.Context) bt.Status {
	if d.child == nil {
		return bt.Success
	}
	if d.child.Tick(ctx) == bt.Running {
		return bt.Running
	}
	d.child.Reset()
	return bt.Success
}

func (d *Succeeder) Reset() {
	if d.child != nil {
		d.child.Reset()
	}
}

func (d *Succeeder) Child() bt.Node {
	return d.child
}

type Failer struct {
	child bt.Node
}

func NewFailer(child bt.Node) *Failer {
	return &Failer{child: child}
}

func (d *Failer) Tick(ctx *bt.Context) bt.Status {
	if d.child == nil {
		return bt.Failure
	}
	if d.child.Tick(ctx) == bt.Running {
		return bt.Running
	}
	d.child.Reset()
	return bt.Failure
}

func (d *Failer) Reset() {
	if d.child != nil {
		d.child.Reset()
	}
}

func (d *Failer) Child() bt.Node {
	return d.child
}

type UntilSuccess struct {
	child bt.Node
}

func NewUntilSuccess(child bt.Node) *UntilSuccess {
	return &UntilSuccess{child: child}
}

func (d *UntilSuccess) Tick(ctx *bt.Context) bt.Status {
	if d.child == nil {
		return bt.Failure
	}
	status := d.child.Tick(ctx)
	if status == bt.Running {
		return bt.Running
	}
	d.child.Reset()
	if status == bt.Success {
		return bt.Success
	}
	return bt.Running
}

func (d *UntilSuccess) Reset() {
	if d.child != nil {
		d.child.Reset()
	}
}

func (d *UntilSuccess) Child() bt.Node {
	return d.child
}

type UntilFailure struct {
	child bt.Node
}

func NewUntilFailure(child bt.Node) *UntilFailure {
	return &UntilFailure{child: child}
}

func (d *UntilFailure) Tick(ctx *bt.Context) bt.Status {
	if d.child == nil {
		return bt.Failure
	}
	status := d.child.Tick(ctx)
	if status == bt.Running {
		return bt.Running
	}
	d.child.Reset()
	if status == bt.Failure {
		return bt.Success
	}
	return bt.Running
}

func (d *UntilFailure) Reset() {
	if d.child != nil {
		d.child.Reset()
	}
}

func (d *UntilFailure) Child() bt.Node {
	return d.child
}
