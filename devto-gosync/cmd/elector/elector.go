package main

import (
	"flag"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/kelein/cookbook/devto-gosync/leader"
)

const logTime = "2006-01-02T15:04:05.000"

var (
	nodeID    = flag.Int("id", os.Getpid(), "node ID")
	etcdAddr  = flag.String("etcd-addr", "localhost:2379", "etcd address")
	etcdAuth  = flag.String("etcd-auth", "", "etcd auth info [user:pass]")
	electName = flag.String("elect-name", "/admin/leader/info", "election name")
)

func init() { initLogger() }

func initLogger() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		if a.Value.Kind() == slog.KindTime {
			return slog.String(a.Key, a.Value.Time().Format(logTime))
		}
		return a
	}

	// * Text Log Format
	multiWriter := io.MultiWriter(os.Stdout)
	logger := slog.New(slog.NewTextHandler(
		multiWriter,
		&slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: replace,
		},
	))
	slog.SetDefault(logger)
}

func main() {
	flag.Parse()
	manager := leader.NewElectManager(*nodeID, *etcdAddr, *etcdAuth, *electName)
	manager.Start()
}
