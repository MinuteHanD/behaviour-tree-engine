package tests

import (
	"testing"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/composite"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/decorator"
	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt/leaf"
)

func TestComplexTree_SurvivingTheWeekend(t *testing.T) {
	board := &MockBlackboard{data: make(map[string]any)}
	ctx := &bt.Context{Blackboard: board}

	board.Set("has_coffee", false)
	board.Set("phone_charged", false)

	hasCoffeeCond := leaf.NewCondition(func(ctx *bt.Context) bool {
		val, ok := ctx.Blackboard.Get("has_coffee")
		return ok && val.(bool)
	})

	phoneChargedCond := leaf.NewCondition(func(ctx *bt.Context) bool {
		val, ok := ctx.Blackboard.Get("phone_charged")
		return ok && val.(bool)
	})

	orderAppAction := leaf.NewAction(func(ctx *bt.Context) bt.Status {
		ctx.Blackboard.Set("has_coffee", true)
		return bt.Success
	})

	tryDeliverySeq := composite.NewSequence(phoneChargedCond, orderAppAction)

	makeCoffeeAction := leaf.NewAction(func(ctx *bt.Context) bt.Status {
		ctx.Blackboard.Set("has_coffee", true)
		return bt.Success
	})

	getCoffeeSelector := composite.NewSelector(tryDeliverySeq, makeCoffeeAction)
	needsCoffee := decorator.NewInverter(hasCoffeeCond)

	rootTree := bt.NewTree(composite.NewSequence(needsCoffee, getCoffeeSelector))

	status := rootTree.Tick(ctx)
	if status != bt.Success {
		t.Errorf("Ecpected success got %v", status)
	}

	hasCoffee, _ := board.Get("has_coffee")
	if !hasCoffee.(bool) {
		t.Errorf("Dont have coffee something broke")
	}

	rootTree.Reset()
	status2 := rootTree.Tick(ctx)
	if status2 != bt.Failure {
		t.Errorf("expected failure got %v", status2)
	}
}

type MockBlackboard struct {
	data map[string]any
}

func (m *MockBlackboard) Get(key string) (any, bool) {
	val, ok := m.data[key]
	return val, ok
}
func (m *MockBlackboard) Set(key string, value any) {
	m.data[key] = value
}
