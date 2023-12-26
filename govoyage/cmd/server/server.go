package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/kelein/cookbook/govoyage/pbgen"
	"github.com/kelein/cookbook/govoyage/service"
)

const (
	contentHeader   = "Content-Type"
	contentTypeGRPC = "application/grpc"
)

var port = flag.Int("port", 8080, "server listen port")

func init() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		// Remove the directory from the source's filename.
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}

	// * Text Log Format
	logger := slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: replace,
			// Level:       slog.LevelDebug,
		},
	))
	slog.SetDefault(logger)
}

func mainV1() {
	flag.Parse()
	slog.Info("server start listen on", "port", *port)

	server := grpc.NewServer()
	tager := service.NewTagService()
	greeter := service.NewGreeterService()
	pbgen.RegisterGreeterServer(server, greeter)
	pbgen.RegisterTagServiceServer(server, tager)

	// * Registe and grpcurl can test it
	reflection.Register(server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Error("server listen failed", "error", err)
	}
	server.Serve(listener)
}

func main() {
	flag.Parse()
	slog.Info("server start listen on", "port", *port)

	if err := RunMultiServer(*port); err != nil {
		slog.Error("server start failed", "error", err)
		os.Exit(1)
	}
}

// RunMultiServer start both http and grpc server on the same port
func RunMultiServer(port int) error {
	addr := fmt.Sprintf(":%d", port)

	httpMux := runHTTPServer()
	grpcServer := runGRPCServer()
	gatewayMux := runGRPCGateway(port)

	httpMux.Handle("/", gatewayMux)
	return http.ListenAndServe(addr, grpcHandlerFunc(grpcServer, httpMux))
}

// runHTTPServer starts a http server via grpc-gateway
func runHTTPServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "HTTP server is running!"}`))
	})
	return mux
}

// runGRPCServer start a grpc server
func runGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	tager := service.NewTagService()
	greeter := service.NewGreeterService()
	pbgen.RegisterGreeterServer(server, greeter)
	pbgen.RegisterTagServiceServer(server, tager)

	// * Registe and grpcurl can test it
	reflection.Register(server)
	return server
}

// runGRPCGateway start a grpc gateway server
func runGRPCGateway(port int) *runtime.ServeMux {
	endpoint := fmt.Sprintf(":%d", port)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}
	pbgen.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	pbgen.RegisterTagServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	return mux
}

func grpcHandlerFunc(grpcServer *grpc.Server, httpServer http.Handler) http.Handler {
	grpcFn := func(w http.ResponseWriter, r *http.Request) {
		ctype := r.Header.Get(contentHeader)
		if r.ProtoMajor == 2 && strings.Contains(ctype, contentTypeGRPC) {
			slog.Info("[GRPC]", "proto", r.Proto, "method", r.Method, "path", r.URL.String())
			grpcServer.ServeHTTP(w, r)
		} else {
			slog.Info("[HTTP]", "proto", r.Proto, "method", r.Method, "path", r.URL.String())
			httpServer.ServeHTTP(w, r)
		}
	}

	return h2c.NewHandler(
		http.HandlerFunc(grpcFn),
		&http2.Server{},
	)
}
