package main

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"goraven/internal/conf"
	logus "goraven/pkg/logger"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gopkg.in/natefinch/lumberjack.v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	logfile := &lumberjack.Logger{
		Filename:   "logs/goraven.log",
		MaxSize:    500,
		MaxBackups: 5,
		MaxAge:     7,
		Compress:   true,
	}
	defer logfile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logfile)

	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		if a.Value.Kind() == slog.KindTime {
			return slog.String(a.Key, a.Value.Time().Format(time.DateTime))
		}
		return a
	}

	opts := &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: replace,
		Level:       slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(multiWriter, opts)
	// handler := slog.NewTextHandler(multiWriter, opts)
	slogger := slog.New(handler)
	slog.SetDefault(slogger)

	klogger := logus.NewSlogger(slogger)

	logger := log.With(klogger,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	slog.Info("server running...")

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
