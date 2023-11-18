package wood

import (
	"testing"

	"github.com/pkg/errors"
)

func TestDecorateF(t *testing.T) {
	setup(InfoLevel, InfoLevel)

	args := []interface{}{
		`%s %s
[%.3fms] [rows:%v] %s`,
		"controlp/db/dbstores/schemadb/store.go:25",
		errors.New("no such table: schemas"),
		3.50775,
		int64(0),
		"X * FROM 'schemas'",
	}

	decorateF(InfoLevel, args, func(format string, args []any) {
		std.Infof(format, args...)
	})
}
