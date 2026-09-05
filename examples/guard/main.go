package main

import (
	"fmt"
	"time"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/blackboard"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/composite"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

type Guard struct {
	Name          string
	PlayerVisible bool
	HeardNoise    bool
}

func main() {
	guard := &Guard{Name: "Ugluk"}
	root := composite.NewReactiveSelector(
		composite.NewSequence(
			leaf.NewCondition(func(ctx *bt.Context) bool {
				return ctx.Agent.(*Guard).PlayerVisible
			}),
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				fmt.Printf("%s attacks\n", ctx.Agent.(*Guard).Name)
				return bt.Success
			}),
		),
		composite.NewSequence(
			leaf.NewCondition(func(ctx *bt.Context) bool {
				return ctx.Agent.(*Guard).HeardNoise
			}),
			leaf.NewAction(func(ctx *bt.Context) bt.Status {
				guard := ctx.Agent.(*Guard)
				fmt.Printf("%s investigates\n", guard.Name)
				guard.HeardNoise = false
				return bt.Success
			}),
		),
		leaf.NewAction(func(ctx *bt.Context) bt.Status {
			fmt.Printf("%s patrols\n", ctx.Agent.(*Guard).Name)
			return bt.Success
		}),
	)

	tree := bt.NewTree(root)
	ctx := bt.NewContext(guard, blackboard.New())
	ctx.DeltaTime = 16 * time.Millisecond

	for frame := 1; frame <= 5; frame++ {
		if frame == 2 {
			guard.HeardNoise = true
		}
		if frame == 4 {
			guard.PlayerVisible = true
		}
		fmt.Printf("frame %d: %s\n", frame, tree.Tick(ctx))
	}
}
