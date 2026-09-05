# Behaviour Tree Engine

[![CI](https://github.com/MinuteHanD/behaviour-tree-engine/actions/workflows/ci.yml/badge.svg)](https://github.com/MinuteHanD/behaviour-tree-engine/actions/workflows/ci.yml)

A dependency-free Go behaviour-tree library for games, simulations, workflows, and autonomous agents.

## Install

```sh
go get github.com/MinuteHanD/behaviour-tree-engine
```

## Quick start

```go
board := blackboard.New()
root := composite.NewReactiveSelector(
	composite.NewSequence(
		leaf.NewCondition(func(ctx *bt.Context) bool {
			return ctx.Agent.(*Guard).PlayerVisible
		}),
		leaf.NewAction(func(ctx *bt.Context) bt.Status {
			return bt.Success
		}),
	),
	leaf.NewAction(func(*bt.Context) bt.Status {
		return bt.Running
	}),
)

tree := bt.NewTree(root)
status := tree.Tick(&bt.Context{
	Agent:      guard,
	DeltaTime: 16 * time.Millisecond,
	Blackboard: board,
})
```

Run the complete example with:

```sh
go run ./examples/guard
```

## Packages

| Package | Purpose |
| --- | --- |
| `pkg/bt` | Core interfaces, tree, context, statuses, and function adapters. |
| `pkg/bt/leaf` | Actions, conditions, fixed statuses, tick waits, duration waits, and named nodes. |
| `pkg/bt/composite` | Memory and reactive sequences/selectors plus threshold parallel control flow. |
| `pkg/bt/decorator` | Inversion, forced results, bounded repeat/retry, until-result loops, and duration timeouts. |
| `pkg/bt/blackboard` | Concurrent in-memory blackboard with typed reads and snapshot support. |

## Semantics

`Sequence` and `Selector` retain their current child while it is running. Use `ReactiveSequence` and `ReactiveSelector` when higher-priority guards must be evaluated on every update.

`Parallel` records terminal child outcomes and never ticks an already-completed child again. Use `NewParallelAll`, `NewParallelAny`, or `NewParallel(required, children...)`.

`Wait(ticks)` returns running for the requested number of ticks and succeeds on the next tick. `WaitDuration` and `Timeout` consume positive `Context.DeltaTime` values. They do not use wall-clock time, so tests and simulations remain deterministic.

`Repeater(count)` completes successfully after running its child `count` times. `Retry(attempts)` succeeds on its first successful attempt or fails after the final failed attempt. Zero repeats succeeds immediately; zero retry attempts fails immediately.

Every composite and decorator resets retained descendant state after producing a terminal result. Trees and built-in nodes accept a nil context; a nil root or nil child fails safely, except for result-forcing decorators whose result is independent of a missing child.

`Tree.Root`, `Tree.Walk`, and `bt.Walk` expose a read-only structural traversal. Composite nodes implement `ChildrenNode`; decorators and named leaves implement `ChildNode`.

## Blackboard

```go
board := blackboard.New()
board.Set("target", "player")

target, ok := blackboard.Get[string](board, "target")
snapshot := board.Snapshot()
board.Delete("target")
```

The blackboard synchronizes access to its map. Values stored inside it remain owned by the caller; synchronize mutable reference values separately.

## Development

```sh
make test
make vet
make race
```

The project targets Go 1.25 or newer.
