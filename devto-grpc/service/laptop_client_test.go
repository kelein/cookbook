package service

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"cookbook/devto-grpc/convert"
	"cookbook/devto-grpc/repo"
	"cookbook/devto-grpc/service"
	"cookbook/devto-grpc/store"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := store.NewMemoryLaptopStore()
	addr := startTestLaptopServer(t, laptopStore)
	client := newTestLaptopClient(t, addr)

	laptop := service.NewLaptop()
	originID := laptop.Id
	req := &repo.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := client.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, originID, res.Id)

	target, err := laptopStore.Find(res.Id)
	requireSameLaptop(t, laptop, target)
}

func requireSameLaptop(t *testing.T, origin *repo.Laptop, target *repo.Laptop) {
	t.Helper()

	originJSON, err := convert.MarshalJSON(origin)
	require.NoError(t, err)

	targetJSON, err := convert.MarshalJSON(target)
	require.NoError(t, err)

	require.Equal(t, originJSON, targetJSON)
}

func startTestLaptopServer(t *testing.T, laptopStore store.LaptopStore) string {
	t.Helper()

	server := grpc.NewServer()
	svc := service.NewLaptopServer(laptopStore)
	repo.RegisterLaptopServiceServer(server, svc)

	lis, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go server.Serve(lis)
	return lis.Addr().String()
}

func newTestLaptopClient(t *testing.T, addr string) repo.LaptopServiceClient {
	t.Helper()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	return repo.NewLaptopServiceClient(conn)
}
