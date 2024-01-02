package tracer

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"sourcegraph.com/sourcegraph/appdash"
	apptracing "sourcegraph.com/sourcegraph/appdash/opentracing"
	"sourcegraph.com/sourcegraph/appdash/traceapp"
)

// STORE memory store for trace collect
var STORE = appdash.NewMemoryStore()

// SetupTracer sets up the tracing system
func SetupTracer() error {
	cfg := config.Configuration{
		ServiceName: "Govoyage",
		Sampler:     &config.SamplerConfig{},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: time.Second * 1,
		},
	}
	tr, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
	)
	opentracing.SetGlobalTracer(tr)
	defer closer.Close()
	return err
}

// SetTracerWithCollector sets up the tracing system with appdash collector
func SetTracerWithCollector(addr string) {
	collector := appdash.NewRemoteCollector(addr)
	tr := apptracing.NewTracer(collector)
	opentracing.SetGlobalTracer(tr)
}

// SetupCollectorServer sets up the Jaeger collector
func SetupCollectorServer() (string, error) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IPv4(0, 0, 0, 0)})
	if err != nil {
		slog.Error("trace collector listener failed", "error", err)
		return "", err
	}
	addr := fmt.Sprintf("0.0.0.0:%d", listener.Addr().(*net.TCPAddr).Port)
	slog.Info("trace collector listen on", "addr", addr)

	// store := appdash.NewMemoryStore()
	server := appdash.NewServer(listener, appdash.NewLocalCollector(STORE))
	go server.Start()
	return addr, nil
}

// SetupCollectorUI sets up the Jaeger collector UI
func SetupCollectorUI(port int) error {
	link := fmt.Sprintf("http://kallen:%d", port)
	appURL, err := url.Parse(link)
	if err != nil {
		slog.Error("parse collector url failed", "error", err)
		return err
	}
	slog.Info("trace collector UI listen", "link", link)

	app, err := traceapp.New(nil, appURL)
	if err != nil {
		slog.Error("start collector UI failed", "error", err)
		return err
	}
	app.Store = STORE
	app.Queryer = STORE

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), app)
		if err != nil {
			slog.Error("collector UI serve failed", "error", err)
		}
	}()
	return nil
}
