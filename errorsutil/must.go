package errorsutil

import (
	"fmt"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}
	return v
}
