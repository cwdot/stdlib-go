package counter

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cwdot/stdlib-go/wood"
)

func New() *CounterMap {
	counter := make(map[string]int)
	return &CounterMap{
		counter: counter,
	}
}

type CounterMap struct {
	mu      sync.Mutex
	counter map[string]int
}

// String debug
func (km *CounterMap) String() string {
	km.mu.Lock()
	defer km.mu.Unlock()

	items := make([]string, 0, len(km.counter))
	for k, v := range km.counter {
		items = append(items, fmt.Sprintf("%s=|%d|", k, v))
	}

	return fmt.Sprintf("CounterMap: %s", strings.Join(items, ", "))
}

// Add saves item; creates path
func (km *CounterMap) Add(key string, increment int) bool {
	wood.Push("Add")
	defer wood.Pop()

	km.mu.Lock()
	defer km.mu.Unlock()

	value, ok := km.counter[key]
	if ok {
		value += increment
	} else {
		value = increment
	}
	km.counter[key] = value
	return true
}

// Get return Item by its id
func (km *CounterMap) Get(key string) (int, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	item, ok := km.counter[key]
	return item, ok
}

func (km *CounterMap) Iterate() map[string]int {
	km.mu.Lock()
	defer km.mu.Unlock()

	clone := make(map[string]int)
	for k, v := range km.counter {
		clone[k] = v
	}
	return clone
}
