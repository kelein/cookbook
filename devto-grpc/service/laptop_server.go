package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cookbook/devto-grpc/repo"
	"cookbook/devto-grpc/store"
)

// LaptopServer provide Laptop service
type LaptopServer struct {
	laptopStore store.LaptopStore
}

// NewLaptopServer crate a new LaptopServer
func NewLaptopServer(laptopStore store.LaptopStore) *LaptopServer {
	return &LaptopServer{laptopStore}
}

// CreateLaptop crate a laptop via unary RPC
func (server *LaptopServer) CreateLaptop(
	ctx context.Context, req *repo.CreateLaptopRequest) (
	*repo.CreateLaptopResponse, error) {

	if ctx.Err() == context.Canceled {
		log.Printf("request canceled")
		return nil, status.Error(codes.Canceled, "request canceled")
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("context deadline exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "context deadline exceeded")
	}

	laptop := req.GetLaptop()
	log.Printf("create laptop request id: %s", laptop.Id)
	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid laptop uuid: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "generate laptop uuid failed: %v", err)
		}
		laptop.Id = id.String()
	}

	err := server.laptopStore.Save(laptop)
	log.Printf("laptop saved result err=%v", err)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, store.ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "save laptop failed: %v", err)
	}
	log.Printf("laptop saved with id: %s", laptop.Id)
	return &repo.CreateLaptopResponse{Id: laptop.Id}, nil
}
