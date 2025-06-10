package creational

import (
	"log/slog"
	"testing"
	"time"
)

func TestBuilder(t *testing.T) {
	pool, err := Builder().DSN("root:root@tcp(127.0.0.1:3306)/test").
		MaxOpenConn(50).MaxIdleConn(10).
		MaxConnLifeTime(time.Hour * 1).Build()
	if err != nil {
		t.Errorf("Builder error: %v", err)
	}
	slog.Info("DB pool build success", "pool", pool)
}
