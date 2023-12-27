package middleware

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

// LoggingInterceptor implements logging UnaryServerInterceptor
func LoggingInterceptor() grpc.UnaryServerInterceptor {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	return logging.UnaryServerInterceptor(
		InterceptorLogger(slog.Default()), opts...,
	)
}

// InterceptorLogger adapts slog logger to interceptor logger
func InterceptorLogger(s *slog.Logger) logging.Logger {
	return logging.LoggerFunc(
		func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			s.Log(ctx, slog.Level(lvl), msg, fields...)
		},
	)
}
