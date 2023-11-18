package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"cookbook/devto-grpc/repo"
	"cookbook/devto-grpc/service"
	"cookbook/devto-grpc/store"
)

var port = flag.Int("port", 0, "the server port")

var imgdir = flag.String("imgdir", "./../tests/imgs", "the image store dir")

func main() {
	flag.Parse()

	server := grpc.NewServer()
	laptopStore := store.NewMemoryLaptopStore()
	imageStore := store.NewDiskImageStore(*imgdir)
	rateStore := store.NewMemoryRateStore()
	svc := service.NewLaptopServer(laptopStore, imageStore, rateStore)
	repo.RegisterLaptopServiceServer(server, svc)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("net listen on [%s] failed: %v", addr, err)
	}
	log.Printf("server start listen on %v", addr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("start server failed: %v", err)
	}
}
