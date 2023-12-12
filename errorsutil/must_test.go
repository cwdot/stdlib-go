package errorsutil

import (
	"testing"

	"github.com/pkg/errors"
)

func TestMustReturnsValueWhenNoError(t *testing.T) {
	expected := "test"
	actual := Must(expected, nil)
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestMustPanicsWhenError(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Must("test", errors.New("test error"))
}
