package creational

import (
	"log/slog"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig(
		WithMaxConns(1000),
		WithTimeout(30),
		WithDebug(),
	)
	slog.Info("server config initialized", "config", cfg)
}
