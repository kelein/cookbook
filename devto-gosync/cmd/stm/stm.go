package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/sync/errgroup"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

const logTime = "2006-01-02T15:04:05.000"

const (
	accountNum  = 5
	origAmount  = 100
	exchangeNum = 999
	accountFmt  = "/accts/%04d"
)

var (
	etcdAddr = flag.String("etcd-addr", "localhost:2379", "etcd address")
	etcdAuth = flag.String("etcd-auth", "", "etcd auth info [user:pass]")
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

	endpoints := strings.Split(*etcdAddr, ",")
	auth := strings.Split(*etcdAuth, ":")
	if len(auth) != 2 {
		slog.Error("etcd auth info invalid")
		os.Exit(1)
	}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:        endpoints,
		Username:         auth[0],
		Password:         auth[1],
		DialTimeout:      time.Second * 3,
		AutoSyncInterval: time.Second * 30,
	})
	if err != nil {
		slog.Error("init etcd client failed", "error", err)
		os.Exit(1)
	}
	defer cli.Close()

	doExchange(cli)

	checkAccount(cli)
}

func doExchange(cli *clientv3.Client) {
	restAccount(cli)

	eg := &errgroup.Group{}
	for range accountNum {
		eg.Go(func(ctx context.Context) error {
			for range exchangeNum {
				_, err := concurrency.NewSwTM(cli, exchange)
				if err != nil {
					slog.Error("stm exchange failed", "error", err)
					return err
				}
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		slog.Error("do exchange failed", "error", err)
		return
	}
}

func exchange(stm concurrency.STM) error {
	from, dest := randN(1, accountNum), randN(1, accountNum)
	if from == dest {
		return nil
	}

	fromK, destK := fmt.Sprintf(accountFmt, from), fmt.Sprintf(accountFmt, dest)
	fromV, destV := stm.Get(fromK), stm.Get(destK)
	fromInt, destInt := 0, 0
	fmt.Sscanf(fromV, "%d", &fromInt)
	fmt.Sscanf(destV, "%d", &destInt)
	slog.Info("account", "from", fromK, "fsum", fromInt, "dest", destK, "dsum", destInt)

	xfer := fromInt / 2
	fromInt, destInt = fromInt-xfer, destInt+xfer
	slog.Info("exchange", "from", fromK, "dest", destK, "xfer", xfer)
	stm.Put(fromK, fmt.Sprintf("%d", fromInt))
	stm.Put(destK, fmt.Sprintf("%d", destInt))
	return nil
}

func checkAccount(cli *clientv3.Client) {
	accts, err := cli.Get(context.TODO(), "/accts/", clientv3.WithPrefix())
	if err != nil {
		slog.Error("search accounts failed", "error", err)
		return
	}

	sum, expectNum := 0, origAmount*accountNum
	for _, kv := range accts.Kvs {
		v := 0
		fmt.Sscanf(string(kv.Value), "%d", &v)
		sum += v
	}
	slog.Info("total accounts checked", "sum", sum, "succeed", sum == expectNum)
}

// restAccount initialize the accounts with origin amount
func restAccount(cli *clientv3.Client) {
	for i := range accountNum {
		k := fmt.Sprintf(accountFmt, i+1)
		_, err := cli.Put(context.TODO(), k, strconv.Itoa(origAmount))
		if err != nil {
			slog.Error("put account failed", "error", err)
			return
		}
	}
	slog.Info("init and reset accounts successfully")
}

func randN(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return min + rand.Intn(max-min+1)
}
