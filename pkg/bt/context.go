package bt

import "time"

type Context struct {
	Agent      any
	DeltaTime  time.Duration
	Blackboard Blackboard
}

type Blackboard interface {
	Get(key string) (any, bool)
	Set(key string, value any)
}

func NewContext(agent any, board Blackboard) *Context {
	return &Context{Agent: agent, Blackboard: board}
}
