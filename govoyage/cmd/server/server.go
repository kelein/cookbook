package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggest/swgui/v5emb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/proxy/grpcproxy"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/kelein/cookbook/govoyage/assets"
	"github.com/kelein/cookbook/govoyage/pbgen"
	"github.com/kelein/cookbook/govoyage/pkg/middleware"
	"github.com/kelein/cookbook/govoyage/pkg/tracer"
	"github.com/kelein/cookbook/govoyage/pkg/version"
	"github.com/kelein/cookbook/govoyage/service"
)

const (
	contentHeader   = "Content-Type"
	contentTypeGRPC = "application/grpc"
)

var (
	v           = flag.Bool("v", false, "show build version")
	vs          = flag.Bool("version", false, "show build version")
	port        = flag.Int("port", 8080, "server listen port")
	etcdAddr    = flag.String("etcd-addr", "127.0.0.1:2379", "etcd server address")
	traceUIPort = flag.Int("trace-ui-port", 8090, "trace collector UI port")
	serverAddr  = fmt.Sprintf("0.0.0.0:%d", *port)
)

var (
	serviceName = version.AppName
	apiDocName  = version.AppName
	apiDocPath  = "/api/v1/docs/"
)

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

	// * Register Prometheus Metrics Collector
	prometheus.MustRegister(version.NewCollector(version.AppName))
}

func main() {
	flag.Parse()
	showVersion()

	slog.Info("server build info", "version", version.Version, "branch",
		version.Branch, "revision", version.Revision, "buildUser",
		version.BuildUser, "buildDate", version.BuildDate)
	slog.Info("server start listen on", "addr", serverAddr)

	// if err := tracer.SetupTracer(); err != nil {
	// 	slog.Error("opentracing setup failed", "error", err)
	// 	os.Exit(1)
	// }

	collectorAddr, err := tracer.SetupCollectorServer()
	if err != nil {
		slog.Error("jaeger collector setup failed", "error", err)
		os.Exit(1)
	}
	tracer.SetTracerWithCollector(collectorAddr)

	if err := tracer.SetupCollectorUI(*traceUIPort); err != nil {
		slog.Error("jaeger collector UI setup failed", "error", err)
		os.Exit(1)
	}

	if err := RunMultiServer(serverAddr); err != nil {
		slog.Error("server start failed", "addr", serverAddr, "error", err)
		os.Exit(1)
	}
}

// mainV1 test unary and stream gRPC server
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

// mainV2 test Swagger UI embedded
func mainV2() {
	// * Register static OpenAPI yaml file
	http.Handle(assets.OpenAPIFilePath, http.FileServer(http.FS(assets.OpenAPIFile)))
	slog.Info("server openAPI file", "path", assets.OpenAPIFilePath)

	// * Register swagger UI and docs
	http.Handle(apiDocPath, v5emb.New(apiDocName, assets.OpenAPIFilePath, apiDocPath))
	slog.Info("server openAPI docs", "path", apiDocPath)

	// * Register index handler
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello Govoyage!"))
	})

	slog.Info("server start listen on", "addr", serverAddr)
	http.ListenAndServe(serverAddr, http.DefaultServeMux)
}

// RunMultiServer start both http and grpc server on the same port
func RunMultiServer(addr string) error {
	httpServer := runHTTPServer()
	grpcServer := runGRPCServer()
	gatewayServer := runGRPCGateway(addr)
	httpServer.Handle("/", gatewayServer)

	// * Register service endpoints
	if err := registerService(addr); err != nil {
		slog.Error("register service failed", "error", err)
		return err
	}
	return http.ListenAndServe(addr, grpcHandlerFunc(grpcServer, httpServer))
}

// registerService register the service into etcd
func registerService(addr string) error {
	etcdcli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{*etcdAddr},
		DialTimeout: time.Second * 60,
	})
	if err != nil {
		slog.Error("connect etcd failed", "addr", etcdAddr, "error", err)
		return err
	}
	defer etcdcli.Close()
	slog.Info("etcd client info", "endpoints", etcdcli.Endpoints())

	logger, err := zap.NewProduction()
	if err != nil {
		slog.Error("initial zap logger failed", "error", err)
		return err
	}
	prefix := fmt.Sprintf("/etcdv3://registry/service/grpc/%s", serviceName)
	grpcproxy.Register(logger, etcdcli, prefix, addr, 60)
	return nil
}

// runHTTPServer starts a http server via grpc-gateway
func runHTTPServer() *http.ServeMux {
	mux := http.NewServeMux()

	// * Register metrics endpoint
	mux.Handle("/metrics", promhttp.Handler())

	// * Register index handler
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(index))
	})
	slog.Info("server home index", "path", "/index")

	// * Register server runtime version info
	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		info, err := json.Marshal(version.Runtime())
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(info)
	})

	// * Register heart beat endpoint
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "HTTP server is running!"}`))
	})

	// * Register static OpenAPI yaml file
	mux.Handle(assets.OpenAPIFilePath, http.FileServer(http.FS(assets.OpenAPIFile)))
	slog.Info("server openAPI file", "path", assets.OpenAPIFilePath)

	// * Register swagger UI and docs
	mux.Handle(apiDocPath, v5emb.New(apiDocName, assets.OpenAPIFilePath, apiDocPath))
	slog.Info("server openAPI docs", "path", apiDocPath)

	// * Register Pprof endpoint
	middleware.NewProfiler().Register(mux)

	return mux
}

// runGRPCServer start a grpc server
func runGRPCServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			middleware.LogInterceptor,
			middleware.NopInterceptor,
			middleware.LoggingInterceptor(),
			middleware.ServerTraceInterceptor(),
		),
	}

	server := grpc.NewServer(opts...)
	tager := service.NewTagService()
	greeter := service.NewGreeterService()
	pbgen.RegisterGreeterServer(server, greeter)
	pbgen.RegisterTagServiceServer(server, tager)

	// * Registe and grpcurl can test it
	reflection.Register(server)
	return server
}

// runGRPCGateway start a grpc gateway server
func runGRPCGateway(addr string) *runtime.ServeMux {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	}
	pbgen.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, addr, opts)
	pbgen.RegisterTagServiceHandlerFromEndpoint(context.Background(), mux, addr, opts)
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
	return h2c.NewHandler(http.HandlerFunc(grpcFn), &http2.Server{})
}

func showVersion() {
	if *v || *vs {
		fmt.Println(version.String())
		os.Exit(0)
	}
}

var index = `<!DOCTYPE html>
<html lang="en"><body>
<h3>Govoyage</h3>
<li><a href="/openapi.yaml">OpenAPI Yaml</a></li>
<li><a href="/api/v1/docs/">OpenAPI Docs<a></li>
</body></html>`
