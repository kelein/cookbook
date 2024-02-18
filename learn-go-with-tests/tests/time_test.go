package tests

import (
	"log/slog"
	"testing"
	"time"
)

func TestTimeDuration(t *testing.T) {
	layout := "2006-01-02 15:04"
	start, _ := time.Parse(layout, "2024-02-06 04:35")
	end, _ := time.Parse(layout, "2024-02-06 15:03")
	du := end.Sub(start)
	slog.Info("elasped time", "duration", du)
}
