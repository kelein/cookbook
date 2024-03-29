package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/kelein/cookbook/devto-grpc/model"
	"github.com/kelein/cookbook/devto-grpc/repo"
	"github.com/kelein/cookbook/devto-grpc/service"
	"github.com/kelein/cookbook/devto-grpc/store"
)

const (
	secretKey     = "secret"
	tokenDuration = time.Minute * 15
)

var port = flag.Int("port", 0, "the server port")

var imgdir = flag.String("imgdir", "./../tests/imgs", "the image store dir")

func main() {
	flag.Parse()

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatalf("load TLS credentials failed: %v", err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
		grpc.Creds(tlsCredentials),
	)

	// interceptor := service.NewAuthInterceptor()
	// server := grpc.NewServer(
	// 	grpc.UnaryInterceptor(interceptor.Unary()),
	// 	grpc.StreamInterceptor(interceptor.Stream()),
	// )

	userStore := store.NewMemoryUserStore()
	if err := seedUser(userStore); err != nil {
		log.Fatalf("seed user failed: %v", err)
	}

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

func unaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Printf("=> unary interceptor: %v", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(server any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("=> stream interceptor: %v", info.FullMethod)
	return handler(server, stream)
}

func createUser(userStore store.UserStore, username, password, role string) error {
	user, err := model.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUser(userStore store.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}

func getAllowRoles() map[string][]string {
	laptopServicePath := "/cookbook.laptopService/"
	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"admin", "user"},
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
	}
	return credentials.NewTLS(config), nil
}
