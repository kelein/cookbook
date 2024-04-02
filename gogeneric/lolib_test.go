package gogeneric

import (
	"log/slog"
	"strings"
	"testing"

	"github.com/samber/lo"
)

func TestLoUsage(t *testing.T) {
	tests := []struct {
		name string
		data []string
	}{
		{"A", []string{""}},
		{"B", []string{"Golang", "Go", "C++", "Go", "Golang"}},
		{"C", []string{"Samuel", "John", "Samuel"}},
		{"D", []string{}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			uniques := lo.Uniq(tc.data)
			slog.Info("Uniq()", "origin", tc.data, "unique", uniques)
			uppers := lo.Map(uniques, func(x string, _ int) string {
				return strings.ToUpper(x)
			})
			slog.Info("Map()", "unique", uniques, "upper", uppers)
			lo.ForEach(uppers, func(x string, i int) {
				slog.Info("ForEach()", "index", i, "item", x)
			})
		})
	}
}
