package middleware

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

// LoggingInterceptor implements logging UnaryServerInterceptor
func LoggingInterceptor() grpc.UnaryServerInterceptor {
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithFieldsFromContext(logTraceID),
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

func logTraceID(ctx context.Context) logging.Fields {
	if span := trace.SpanContextFromContext(ctx); span.IsSampled() {
		return logging.Fields{"traceID", span.TraceID().String()}
	}
	return nil
}
