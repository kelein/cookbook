package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cookbook/devto-grpc/repo"
	"cookbook/devto-grpc/store"
)

// maxChunkSizeBytes 1MB
const maxImageSizeBytes = 1 << 20

// LaptopServer provide Laptop service
type LaptopServer struct {
	laptopStore store.LaptopStore
	imageStore  store.ImageStore
	rateStore   store.RateStore

	// UnimplementedLaptopServiceServer must be embedded
	// to have forward compatible implementations.
	repo.UnimplementedLaptopServiceServer
}

// NewLaptopServer crate a new LaptopServer
func NewLaptopServer(laptopStore store.LaptopStore, imageStore store.ImageStore, rateStore store.RateStore) *LaptopServer {
	return &LaptopServer{
		laptopStore: laptopStore,
		imageStore:  imageStore,
		rateStore:   rateStore,
	}
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
	if err != nil {
		log.Printf("laptop saved result err: %v", err)
		code := codes.Internal
		if errors.Is(err, store.ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "save laptop failed: %v", err)
	}
	log.Printf("laptop saved with id: %s", laptop.Id)
	return &repo.CreateLaptopResponse{Id: laptop.Id}, nil
}

// SearchLaptop query laptops via a stream RPC
func (server *LaptopServer) SearchLaptop(req *repo.SearchLaptopRequest, stream repo.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("search laptop request with filter: %v", filter)
	got := func(laptop *repo.Laptop) error {
		res := &repo.SearchLaptopResponse{Laptop: laptop}
		if err := stream.Send(res); err != nil {
			return fmt.Errorf("server stream error: %w", err)
		}
		log.Printf("search request found laptop: %s", laptop.Id)
		return nil
	}
	err := server.laptopStore.Search(stream.Context(), filter, got)
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	return nil
}

// UploadImage upload image file via stream RPC
func (server *LaptopServer) UploadImage(stream repo.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		serr := status.Errorf(codes.Unknown, "revc image failed: %v", err)
		log.Print(serr)
		return serr
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("uploadImage receive laptopID: %v, imageType: %s", laptopID, imageType)
	laptop, err := server.laptopStore.Find(laptopID)
	if err != nil {
		serr := status.Errorf(codes.Internal, "find laptop error: %v", err)
		log.Print(serr)
		return serr
	}
	if laptop == nil {
		serr := status.Errorf(codes.InvalidArgument, "laptop not found: %v", err)
		log.Print(serr)
		return serr
	}

	imageSize := 0
	imageData := bytes.Buffer{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("stream receive End of file")
			break
		}
		if err != nil {
			serr := status.Errorf(codes.Unknown, "receive chunk data failed: %v", err)
			log.Print(serr)
			return serr
		}

		chunk := req.GetChunkData()
		size := len(chunk)
		log.Printf("received chunk data size: %d", size)

		imageSize += size
		if imageSize > maxImageSizeBytes {
			serr := status.Errorf(codes.InvalidArgument, "chunk data exceeded max value: %v", err)
			log.Print(serr)
			return serr
		}
		if _, err := imageData.Write(chunk); err != nil {
			serr := status.Errorf(codes.Internal, "write chunk data failed: %v", err)
			log.Print(serr)
			return serr
		}
	}

	imageID, err := server.imageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		serr := status.Errorf(codes.Internal, "store image failed: %v", err)
		log.Print(serr)
		return serr
	}
	res := &repo.UploadImageResonse{
		Id:   imageID,
		Size: uint32(imageSize),
	}
	if err := stream.SendAndClose(res); err != nil {
		serr := status.Errorf(codes.Unknown, "send response failed: %v", err)
		log.Print(serr)
		return serr
	}
	log.Printf("saved image id: %s, size: %d", imageID, imageSize)
	return nil
}

// RateLaptop rate laptop score via streaming
func (server *LaptopServer) RateLaptop(stream repo.LaptopService_RateLaptopServer) error {
	for {

	}
	return nil
}
