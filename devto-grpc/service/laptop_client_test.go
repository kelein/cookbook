package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/kelein/cookbook/devto-grpc/convert"
	"github.com/kelein/cookbook/devto-grpc/repo"
	"github.com/kelein/cookbook/devto-grpc/store"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := store.NewMemoryLaptopStore()
	imgStore := store.NewDiskImageStore("../test/imgs")
	_, addr := startTestLaptopServer(t, laptopStore, imgStore)
	client := newTestLaptopClient(t, addr)

	laptop := NewLaptop()
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

func startTestLaptopServer(t *testing.T, laptopStore store.LaptopStore, imgStore store.ImageStore) (*LaptopServer, string) {
	t.Helper()

	server := grpc.NewServer()
	rateStore := store.NewMemoryRateStore()
	svc := NewLaptopServer(laptopStore, imgStore, rateStore)
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

	laptopStore := store.NewMemoryLaptopStore()
	imgStore := store.NewDiskImageStore("../test/imgs")
	targetIds := genTestDataForSearch(t, laptopStore)

	_, addr := startTestLaptopServer(t, laptopStore, imgStore)
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

	n0 := NewLaptop()
	n0.PriceUsd = 2500

	n1 := NewLaptop()
	n1.Cpu.Cores = 2

	n2 := NewLaptop()
	n2.Cpu.MinGhz = 2.0

	n3 := NewLaptop()
	n3.Ram = &repo.Memory{Value: 4096, Unit: repo.Memory_MEGABYTE}

	n4 := NewLaptop()
	n4.PriceUsd = 1999
	n4.Cpu.Cores = 4
	n4.Cpu.MinGhz = 2.5
	n4.Cpu.MaxGhz = 4.5
	n4.Ram = &repo.Memory{Value: 16, Unit: repo.Memory_GIGABYTE}
	targetIds[n4.Id] = true

	n5 := NewLaptop()
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

func TestClientUploadImage(t *testing.T) {
	t.Parallel()
	folder := "../tests/imgs"
	imgStore := store.NewDiskImageStore(folder)
	laptopStore := store.NewMemoryLaptopStore()
	laptop := NewLaptop()
	err := laptopStore.Save(laptop)
	require.NoError(t, err)

	_, addr := startTestLaptopServer(t, laptopStore, imgStore)
	t.Logf("test server addr: %q", addr)
	client := newTestLaptopClient(t, addr)

	imgPath := fmt.Sprintf("%s/laptop.png", folder)
	f, err := os.Open(imgPath)
	defer f.Close()
	require.NoError(t, err)
	stream, err := client.UploadImage(context.Background())
	require.NoError(t, err)

	imageType := filepath.Ext(imgPath)
	t.Logf("imageType: %v", imageType)
	req := &repo.UploadImageRequest{
		Data: &repo.UploadImageRequest_Info{
			Info: &repo.ImageInfo{
				LaptopId:  laptop.GetId(),
				ImageType: imageType,
			},
		},
	}
	err = stream.Send(req)
	require.NoError(t, err)

	reader := bufio.NewReader(f)
	buffer := make([]byte, 1024)
	size := 0
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			t.Logf("End of File: %v", f.Name())
			break
		}
		require.NoError(t, err)
		size += n

		req := &repo.UploadImageRequest{
			Data: &repo.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		err = stream.Send(req)
		require.NoError(t, err)
	}

	res, err := stream.CloseAndRecv()
	require.NoError(t, err)
	require.NotZero(t, res.GetId())
	require.EqualValues(t, size, res.GetSize())
	savedImgPath := fmt.Sprintf("%s/%s%s", folder, res.GetId(), imageType)
	require.FileExists(t, savedImgPath)
}
