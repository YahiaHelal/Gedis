package dispatcher

import (
	"sync"
)

type Dispatcher[T comparable, V any] struct {
	handlers map[T]V
	mu       *sync.RWMutex
}

func NewDispatcher[T comparable, V any]() *Dispatcher[T, V] {
	return &Dispatcher[T, V]{handlers: map[T]V{}, mu: &sync.RWMutex{}}
}

func (d *Dispatcher[T, V]) Register(key T, handler V) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[key] = handler
}

func (d *Dispatcher[T, V]) Dispatch(key T) (V, bool) {
	var zero V
	d.mu.RLock()
	defer d.mu.RUnlock()
	if handler, ok := d.handlers[key]; ok {
		return handler, ok
	}
	return zero, false
}

func (d *Dispatcher[T, V]) HasHandler(key T) bool {
	_, ok := d.handlers[key]
	return ok
}
