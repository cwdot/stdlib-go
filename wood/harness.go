package wood

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cwdot/stdlib-go/color"
)

func New() *TestHarness {
	return &TestHarness{SB: new(strings.Builder)}
}

type TestHarness struct {
	SB *strings.Builder
}

func (th *TestHarness) Pop() []Entry {
	text := th.SB.String()
	th.SB.Reset()
	fmt.Printf("pop:  %s\n", text)

	lines := strings.Split(text, "\n")

	entries := make([]Entry, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		var ll Entry
		if err := json.Unmarshal([]byte(line), &ll); err != nil {
			panic(err)
		}

		lvl, err := ParseLevel(ll.LevelText)
		if err != nil {
			panic(err)
		}
		ll.Level = lvl

		// remove any space prefixes first then separate
		trimmed := strings.TrimLeft(ll.Msg, " ")
		before, after, found := strings.Cut(trimmed, " ")
		if found {
			ll.Component = color.Strip(before)
			ll.Msg = strings.TrimLeft(after, " ")
		}
		fmt.Printf("set ll: %s %s\n", ll.Component, ll.Msg)
		entries = append(entries, ll)
	}

	return entries
}

func (th *TestHarness) Assert(t *testing.T, entries []Entry, idx int, level Level, component string, msg string) {
	e := entries[idx]
	assert.Equal(t, level, e.Level)
	assert.Equal(t, component, e.Component)
	assert.Equal(t, msg, e.Msg)
}

type Entry struct {
	LevelText string    `json:"level"`
	Msg       string    `json:"msg"`
	Time      time.Time `json:"time"`

	Level     Level  `json:"-"`
	Component string `json:"-"`
}
