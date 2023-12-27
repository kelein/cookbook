package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	grpcw "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/kelein/cookbook/govoyage/pbgen"
)

var addr = flag.String("address", ":8080", "the server address")

func init() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}

	// * JSON Log Format
	// logger := slog.New(slog.NewJSONHandler(
	// 	os.Stdout, &slog.HandlerOptions{AddSource: true},
	// ))

	// * Text Log Format
	logger := slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: replace,
		},
	))

	slog.SetDefault(logger)
}

func main() {
	flag.Parse()

	// * Client Interceptor Options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),

		grpc.WithUnaryInterceptor(grpcw.ChainUnaryClient(
			// * Retry Interceptor Option
			retry.UnaryClientInterceptor(
				retry.WithMax(2), retry.WithCodes(
					codes.Unknown,
					codes.Internal,
					codes.DeadlineExceeded,
				),
			),
		)),
	}

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		slog.Error("connect server failed", "error", err)
	}
	defer conn.Close()

	client := pbgen.NewGreeterClient(conn)

	// * Simple Unary Request
	multiGreetProcess(client, 20)

	// * Server Stream Request
	mockGreetStream(client, &pbgen.GreetRequest{Name: gofakeit.Name()})

	// * Client Stream Request
	mockGreetRecord(client, &pbgen.GreetRequest{Name: gofakeit.Name()})
}

// mockGreetRecord mock client stream request
func mockGreetRecord(client pbgen.GreeterClient, req *pbgen.GreetRequest) error {
	stream, err := client.GreetRecord(context.Background())
	if err != nil {
		slog.Error("client greet record failed", "error", err)
		return err
	}
	for i := 0; i < 10; i++ {
		if err := stream.Send(req); err != nil {
			slog.Error("client send stream request failed", "error", err)
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		slog.Error("client close stream request failed", "error", err)
		return err
	}
	slog.Info("client received stream result", "message", resp.GetMessage())
	return nil
}

// mockGreetStream mock server stream response
func mockGreetStream(client pbgen.GreeterClient, req *pbgen.GreetRequest) error {
	stream, err := client.GreetStream(context.Background(), req)
	if err != nil {
		slog.Error("client stream request failed", "error", err)
		return err
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			slog.Error("client stream receive failed", "error", err)
			return err
		}
		slog.Info("client stream receive", "message", resp.GetMessage())
	}
	return nil
}

// multiGreetProcess send greet request with multiple goroutines
func multiGreetProcess(client pbgen.GreeterClient, num int) error {
	eg := errgroup.Group{}
	if num <= 0 {
		slog.Error("invalid concurrency", "num", num)
	}

	for i := 0; i < num; i++ {
		idx := fmt.Sprintf("%04d", i+1)
		time.Sleep(time.Second * 1 / 3)
		eg.Go(func() error {
			slog.Info("start greeter routine", "index", idx, "name", gofakeit.Name())
			return mockGreet(client, idx)
		})
	}

	return eg.Wait()
}

// mockGreet mock a simple unary request
func mockGreet(client pbgen.GreeterClient, name string) error {
	resp, err := client.Greet(context.Background(), &pbgen.GreetRequest{Name: name})
	if err != nil {
		slog.Error("client send request failed", "error", err)
		return err
	}
	slog.Info("client got response", "message", resp.Message)
	return nil
}
