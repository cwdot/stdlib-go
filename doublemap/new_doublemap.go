package doublemap

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cwdot/go-stdlib/wood"
)

func New() *DoubleMap {
	counter := make(map[string]map[string]string)
	return &DoubleMap{
		m: counter,
	}
}

type DoubleMap struct {
	mu sync.Mutex
	m  map[string]map[string]string
}

// String debug
func (km *DoubleMap) String() string {
	km.mu.Lock()
	defer km.mu.Unlock()

	items := make([]string, 0, len(km.m))
	for k, v := range km.m {
		items = append(items, fmt.Sprintf("%s=|%d|", k, v))
	}

	return fmt.Sprintf("DoubleMap: %s", strings.Join(items, ", "))
}

// Add saves item; creates path
func (km *DoubleMap) Add(parent string, key string, value string) bool {
	wood.Push("Add")
	defer wood.Pop()

	km.mu.Lock()
	defer km.mu.Unlock()

	if inner, ok := km.m[parent]; ok {
		inner[key] = value
	} else {
		km.m[parent] = make(map[string]string)
		km.m[parent][key] = value
	}
	return true
}

// Get return Item by its id
func (km *DoubleMap) Get(parent string, key string) (string, bool, bool) {
	km.mu.Lock()
	defer km.mu.Unlock()

	if outer, ok := km.m[parent]; ok {
		inner, ok := outer[key]
		return inner, true, ok
	}
	return "", false, false
}
