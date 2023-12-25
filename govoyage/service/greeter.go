package service

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/kelein/cookbook/govoyage/pbgen"
)

// GreeterService provides a Greeter service
type GreeterService struct {
	pbgen.UnimplementedGreeterServer
}

// NewGreeterService creates a new Greeter service
func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

// Greet gives a Greeter reply
func (s *GreeterService) Greet(ctx context.Context, req *pbgen.GreetRequest) (*pbgen.GreetResponse, error) {
	return &pbgen.GreetResponse{
		Message: "Welcome to go!",
	}, nil
}

// GreetStream gives a stream Greeter reply
func (s *GreeterService) GreetStream(req *pbgen.GreetRequest, stream pbgen.Greeter_GreetStreamServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&pbgen.GreetResponse{
			Message: fmt.Sprintf("Welcome to go, %s", req.GetName()),
		})
		if err != nil {
			slog.Error("greet stream failed", "error", err)
			return err
		}
	}
	return nil
}

// GreetRecord send greet reply from client stream
func (s *GreeterService) GreetRecord(stream pbgen.Greeter_GreetRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			msg := &pbgen.GreetResponse{
				Message: "client steam reply",
			}
			return stream.SendAndClose(msg)
		}
		slog.Info("server received result", "name", resp.Name)
	}
}
