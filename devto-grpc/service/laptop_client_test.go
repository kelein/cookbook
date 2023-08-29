package service

import (
	"context"
	"io"
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
	_, addr := startTestLaptopServer(t, laptopStore)
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

func startTestLaptopServer(t *testing.T, laptopStore store.LaptopStore) (*service.LaptopServer, string) {
	t.Helper()

	server := grpc.NewServer()
	svc := service.NewLaptopServer(laptopStore)
	repo.RegisterLaptopServiceServer(server, svc)

	lis, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go server.Serve(lis)
	return svc, lis.Addr().String()
}

func newTestLaptopClient(t *testing.T, addr string) repo.LaptopServiceClient {
	t.Helper()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	require.NoError(t, err)
	return repo.NewLaptopServiceClient(conn)
}

func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()

	store := store.NewMemoryLaptopStore()
	targetIds := genTestDataForSearch(t, store)

	_, addr := startTestLaptopServer(t, store)
	client := newTestLaptopClient(t, addr)

	filter := &repo.Filter{
		MaxPriceUsed: 2000,
		MinCpuCores:  4,
		MinCpuGhz:    2.2,
		MinRam: &repo.Memory{
			Value: 8,
			Unit:  repo.Memory_GIGABYTE,
		},
	}
	req := &repo.SearchLaptopRequest{Filter: filter}
	stream, err := client.SearchLaptop(context.Background(), req)
	require.NoError(t, err)

	found := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		require.Contains(t, targetIds, res.GetLaptop().GetId())
		found++
	}
	require.Equal(t, len(targetIds), found)
}

func genTestDataForSearch(t *testing.T, store store.LaptopStore) map[string]bool {
	t.Helper()
	targetIds := make(map[string]bool)

	n0 := service.NewLaptop()
	n0.PriceUsd = 2500

	n1 := service.NewLaptop()
	n1.Cpu.Cores = 2

	n2 := service.NewLaptop()
	n2.Cpu.MinGhz = 2.0

	n3 := service.NewLaptop()
	n3.Ram = &repo.Memory{Value: 4096, Unit: repo.Memory_MEGABYTE}

	n4 := service.NewLaptop()
	n4.PriceUsd = 1999
	n4.Cpu.Cores = 4
	n4.Cpu.MinGhz = 2.5
	n4.Cpu.MaxGhz = 4.5
	n4.Ram = &repo.Memory{Value: 16, Unit: repo.Memory_GIGABYTE}
	targetIds[n4.Id] = true

	n5 := service.NewLaptop()
	n5.PriceUsd = 2000
	n5.Cpu.Cores = 6
	n5.Cpu.MinGhz = 2.8
	n5.Cpu.MaxGhz = 4.8
	n5.Ram = &repo.Memory{Value: 64, Unit: repo.Memory_GIGABYTE}
	targetIds[n5.Id] = true

	laptops := []*repo.Laptop{n0, n1, n2, n3, n4, n5}
	for _, laptop := range laptops {
		err := store.Save(laptop)
		require.NoError(t, err)
	}
	return targetIds
}
