package middleware

import (
	"context"
	"log/slog"

	"github.com/kelein/cookbook/govoyage/pkg/metatext"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// ServerTraceInterceptor implements server tracing
func ServerTraceInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		parentSpan, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, metatext.MetadataTextMap{MD: md})
		if err != nil {
			slog.Error("tracer extract metadata failed", "error", err)
		}

		spanOpts := []opentracing.StartSpanOption{
			opentracing.Tag{Key: string(ext.Component), Value: "GPRC"},
			ext.SpanKindRPCServer,
			ext.RPCServerOption(parentSpan),
		}
		span := opentracing.GlobalTracer().StartSpan(info.FullMethod, spanOpts...)
		defer span.Finish()

		ctx = opentracing.ContextWithSpan(ctx, span)
		return handler(ctx, req)
	}
}

// ClientTraceInterceptor implements client tracing interceptor
func ClientTraceInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, conn *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		var parentCtx opentracing.SpanContext
		var spanOpts []opentracing.StartSpanOption

		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentCtx = parentSpan.Context()
			spanOpts = append(spanOpts, opentracing.ChildOf(parentCtx))
		}
		spanOpts = append(spanOpts,
			[]opentracing.StartSpanOption{
				opentracing.Tag{Key: string(ext.Component), Value: "GPRC"},
				ext.SpanKindRPCClient,
			}...,
		)
		span := opentracing.GlobalTracer().StartSpan(method, spanOpts...)
		defer span.Finish()

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}

		err := opentracing.GlobalTracer().Inject(span.Context(), opentracing.TextMap, metatext.MetadataTextMap{MD: md})
		if err != nil {
			slog.Error("global tracer inject metadata failed", "error", err)
		}

		newCtx := opentracing.ContextWithSpan(ctx, span)
		return invoker(newCtx, method, req, reply, conn, opts...)
	}
}
