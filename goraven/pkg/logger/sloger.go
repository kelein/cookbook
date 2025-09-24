package logus

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

// Slogger 是 slog.Logger 的包装器，实现了 kratos/log.Logger 接口
type Slogger struct {
	logger *slog.Logger
}

// NewSlogger 创建一个新的 SlogLogger 实例
// 可以通过传入不同的 slog.Handler 来控制输出格式（JSON 或 Text）和行为
func NewSlogger(logger *slog.Logger) *Slogger {
	if logger == nil {
		// 默认使用 JSON Handler，输出到标准错误
		logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	}
	return &Slogger{logger: logger}
}

// Log 实现 kratos/log.Logger 接口的 Log 方法
// 这是最关键的方法，它将 Kratos 的日志调用转换为 slog 的日志记录
func (s *Slogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals)%2 != 0 {
		// 如果 keyvals 不是成对的，slog 可能会处理不了，这里可以加个默认 key
		keyvals = append([]interface{}{"msg"}, keyvals...)
	}

	// 将 Kratos 的日志级别转换为 slog 的日志级别
	slogLevel := toSlogLevel(level)

	// 使用 slog.LogAttrs 可以避免装箱开销，性能更好
	// 这里先将平铺的 keyvals 转换为 slog.Attr
	attrs := make([]slog.Attr, 0, len(keyvals)/2)
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			continue // 或者处理错误
		}
		attrs = append(attrs, slog.Any(key, keyvals[i+1]))
	}

	args := make([]any, 0, len(attrs))
	for _, attr := range attrs {
		args = append(args, attr)
	}

	s.logger.Log(context.Background(), slogLevel, "", slog.Group("", args...))
	return nil
}

// toSlogLevel 将 Kratos 日志级别转换为 slog 日志级别
func toSlogLevel(level log.Level) slog.Level {
	if v, ok := slogLevelMap[level]; ok {
		return v
	}
	return slog.LevelInfo
}

var slogLevelMap = map[log.Level]slog.Level{
	log.LevelDebug: slog.LevelDebug,
	log.LevelInfo:  slog.LevelInfo,
	log.LevelWarn:  slog.LevelWarn,
	log.LevelError: slog.LevelError,
	log.LevelFatal: slog.LevelError,
}

// 为 Fatal 级别提供特殊处理
// Kratos 的 Helper 在调用 Fatal 方法时，会先调用 Log 方法，然后调用 os.Exit(1)
// 因此我们需要确保在 Log 方法中正确处理 LevelFatal
