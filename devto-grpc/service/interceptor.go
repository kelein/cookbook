package service

import (
	"context"
	"cookbook/devto-grpc/pkg/jwtool"
	"log"

	"google.golang.org/grpc"
)

// AuthInterceptor interceptor for authentication
type AuthInterceptor struct {
	jwtManager *jwtool.Manager
	allowRoles map[string]string
}

// NewAuthInterceptor creates a new AuthInterceptor
func NewAuthInterceptor(manager *jwtool.Manager, roles map[string]string) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager: manager,
		allowRoles: roles,
	}
}

// Unary implements Unary Interface
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		log.Printf("=> unary interceptor: %v", info.FullMethod)

		// TODO: implement authentication

		return handler(ctx, req)
	}
}
