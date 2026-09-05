package blackboard

import (
	"sync"

	"github.com/MinuteHanD/behaviour-tree-engine/pkg/bt"
)

type Blackboard struct {
	mu   sync.RWMutex
	data map[string]any
}

func New() *Blackboard {
	return &Blackboard{data: make(map[string]any)}
}

func (b *Blackboard) Get(key string) (any, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	value, ok := b.data[key]
	return value, ok
}

func (b *Blackboard) Set(key string, value any) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.data == nil {
		b.data = make(map[string]any)
	}
	b.data[key] = value
}

func (b *Blackboard) Delete(key string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	delete(b.data, key)
}

func (b *Blackboard) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.data = make(map[string]any)
}

func (b *Blackboard) Len() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return len(b.data)
}

func (b *Blackboard) Snapshot() map[string]any {
	b.mu.RLock()
	defer b.mu.RUnlock()

	values := make(map[string]any, len(b.data))
	for key, value := range b.data {
		values[key] = value
	}
	return values
}

func Get[T any](board bt.Blackboard, key string) (T, bool) {
	var zero T
	if board == nil {
		return zero, false
	}

	value, ok := board.Get(key)
	if !ok {
		return zero, false
	}

	typed, ok := value.(T)
	return typed, ok
}
