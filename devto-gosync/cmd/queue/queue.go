package main

import (
	"bufio"
	"flag"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	recipe "go.etcd.io/etcd/client/v3/experimental/recipes"
)

const logTime = "2006-01-02T15:04:05.000"

const (
	cmdPush = "push"
	cmdPop  = "pop"
	cmdQuit = "quit"
	cmdExit = "exit"
)

var (
	queueName = flag.String("name", "/distriq", "etcd queue name")
	etcdAddr  = flag.String("etcd-addr", "localhost:2379", "etcd address")
	etcdAuth  = flag.String("etcd-auth", "", "etcd auth info [user:pass]")
	priority  = flag.Bool("priority", false, "whether enable queue priority")
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

	slog.Info("etcd queue started", "addr", *etcdAddr, "name", *queueName)

	endpoints := strings.Split(*etcdAddr, ",")
	auth := strings.Split(*etcdAuth, ":")
	if len(auth) != 2 {
		slog.Error("etcd auth info invalid")
		os.Exit(1)
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		Username:  auth[0],
		Password:  auth[1],
	})
	if err != nil {
		slog.Error("init etcd client failed", "error", err)
		os.Exit(1)
	}
	defer cli.Close()

	if *priority {
		consoleScannerWithPriority(cli)
		return
	}
	consoleScanner(cli)
}

func consoleScanner(cli *clientv3.Client) {
	Q := recipe.NewQueue(cli, *queueName)
	slog.Info("etcd queue created", "name", Q)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		action := scanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case cmdPush:
			if len(items) != 2 {
				slog.Error("must set value to push")
				continue
			}
			if err := Q.Enqueue(items[1]); err != nil {
				slog.Error("enqueue failed", "error", err)
				continue
			}
			slog.Info("enqueue item", "value", items[1])
		case cmdPop:
			value, err := Q.Dequeue()
			if err != nil {
				slog.Error("dequeue failed", "error", err)
				continue
			}
			slog.Info("dequeue item", "value", value)
		case cmdQuit, cmdExit:
			return
		default:
			slog.Error("unknown command action")
		}
	}
}

func consoleScannerWithPriority(cli *clientv3.Client) {
	Q := recipe.NewPriorityQueue(cli, *queueName)
	slog.Info("etcd priority queue created", "name", Q)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		action := scanner.Text()
		items := strings.Split(action, " ")
		switch items[0] {
		case cmdPush:
			if len(items) != 3 {
				slog.Error("must set value and priority to push")
				continue
			}
			prio, err := strconv.Atoi(items[2])
			if err != nil {
				slog.Error("invalid priority value", "error", err)
				continue
			}
			if err := Q.Enqueue(items[1], uint16(prio)); err != nil {
				slog.Error("enqueue failed", "error", err)
				continue
			}
			slog.Info("enqueue item", "value", items[1])
		case cmdPop:
			value, err := Q.Dequeue()
			if err != nil {
				slog.Error("dequeue failed", "error", err)
				continue
			}
			slog.Info("dequeue item", "value", value)
		case cmdQuit, cmdExit:
			return
		default:
			slog.Error("unknown command action")
		}
	}
}
