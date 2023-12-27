package middleware

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

// LogInterceptor logs grpc request records
func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	slog.Debug("[GRPC-BEFORE]", "path", info.FullMethod)
	resp, err := handler(ctx, req)
	slog.Info("[GRPC] DONE", "path", info.FullMethod, slog.Duration("duration", time.Since(start)))
	return resp, err
}

// NopInterceptor logs grpc request records
func NopInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	return handler(ctx, req)
}
