# Behaviour Tree Engine

[![CI](https://github.com/MinuteHanD/behaviour-tree-engine/actions/workflows/ci.yml/badge.svg)](https://github.com/MinuteHanD/behaviour-tree-engine/actions/workflows/ci.yml)

A dependency-free Go behaviour-tree library for game engines and autonomous agents.

## Install

```sh
go get github.com/MinuteHanD/behaviour-tree-engine
```

## Packages

| Package | Purpose |
| --- | --- |
| `pkg/bt` | Core interfaces, tree, context, statuses, and function adapters. |
| `pkg/bt/leaf` | Actions, conditions, fixed statuses, tick waits, duration waits, and named nodes. |
| `pkg/bt/composite` | Memory and reactive sequences/selectors plus threshold parallel control flow. |
| `pkg/bt/decorator` | Inversion, forced results, bounded repeat/retry, until-result loops, and duration timeouts. |
| `pkg/bt/blackboard` | Concurrent in-memory blackboard with typed reads and snapshot support. |

## Examples

Runnable scenarios demonstrating emergent AI behaviour can be found in the `examples/` directory:

- `examples/heist`: A dual-agent simulation featuring a Thief and a Cop interacting via a shared blackboard.
- `examples/guard`: A simple reactive patrol/attack guard.

```sh
go run ./examples/heist/main.go
```

## Development

```sh
make test
make vet
make race
```

The project targets Go 1.25 or newer.
