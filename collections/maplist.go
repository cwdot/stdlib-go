package collections

import (
	"fmt"
	"strings"
	"sync"
)

func NewMapList[T any]() *MapList[T] {
	counter := make(map[string][]T)
	return &MapList[T]{
		m: counter,
	}
}

type MapList[T any] struct {
	mu sync.Mutex
	m  map[string][]T
}

// String debug
func (km *MapList[T]) String() string {
	km.mu.Lock()
	defer km.mu.Unlock()

	items := make([]string, 0, len(km.m))
	for k, v := range km.m {
		items = append(items, fmt.Sprintf("%s=%v", k, v))
	}

	return fmt.Sprintf("MapList: %s", strings.Join(items, ", "))
}

// Add saves item; creates path
func (km *MapList[T]) Add(key string, value T) bool {
	km.mu.Lock()
	defer km.mu.Unlock()

	if _, ok := km.m[key]; !ok {
		km.m[key] = make([]T, 0, 5)
	}
	km.m[key] = append(km.m[key], value)

	return true
}

// Get return Item by its id
func (km *MapList[T]) Get(parent string, index int) (T, bool, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	if outer, ok := km.m[parent]; ok {
		return outer[index], true, ok
	}
	var result T
	return result, false, false
}

// GetList return Item by its id
func (km *MapList[T]) GetList(key string) ([]T, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	if outer, ok := km.m[key]; ok {
		return outer, true
	}
	return nil, false
}

// Size
func (km *MapList[T]) Size() int {
	km.mu.Lock()
	defer km.mu.Unlock()

	return len(km.m)
}
