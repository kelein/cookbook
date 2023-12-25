package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kelein/cookbook/govoyage/pbgen"
	"github.com/kelein/cookbook/govoyage/service"
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
		},
	))
	slog.SetDefault(logger)
}

func main() {
	flag.Parse()
	slog.Info("server start listen on", "port", *port)

	server := grpc.NewServer()
	greeter := service.NewGreeterService()
	pbgen.RegisterGreeterServer(server, greeter)

	// * Registe and grpcurl can test it
	reflection.Register(server)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		slog.Error("server listen failed", "error", err)
	}
	server.Serve(listener)
}
