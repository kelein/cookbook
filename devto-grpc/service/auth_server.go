package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cookbook/devto-grpc/pkg/jwtool"
	"cookbook/devto-grpc/repo"
	"cookbook/devto-grpc/store"
)

// AuthServer provide auth service
type AuthServer struct {
	userStore  store.UserStore
	jwtManager *jwtool.Manager
}

// NewAuthServer creates a new AuthServer
func NewAuthServer(userStore store.UserStore, jwtManager *jwtool.Manager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}

// Login loged user into AuthServer
func (server *AuthServer) Login(ctx context.Context, req *repo.LoginRequest) (*repo.LoginResponse, error) {
	user, err := server.userStore.Find(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "find user error: %v", err)
	}
	if user == nil || !user.Authed(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username or password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "generate token failed: %v", err)
	}
	return &repo.LoginResponse{
		AccessToken: token,
	}, nil
}
