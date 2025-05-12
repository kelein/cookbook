package leader

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

const (
	cmdElect    = "elect"
	cmdProclaim = "proclaim"
	cmdResign   = "resign"
	cmdWatch    = "watch"
	cmdQuery    = "query"
	cmdRevision = "revision"
)

// ElectManager manage leader election
type ElectManager struct {
	nodeID    int
	round     int
	etcdAddr  string
	etcdAuth  string
	electName string
}

// NewElectManager create a new ElectManager instance
func NewElectManager(nodeID int, etcdAddr, etcdAuth, electName string) *ElectManager {
	return &ElectManager{
		nodeID:    nodeID,
		etcdAddr:  etcdAddr,
		etcdAuth:  etcdAuth,
		electName: electName,
	}
}

// Start start a node with leader election
func (em *ElectManager) Start() error {
	slog.Info("leader manger started", "node", em.nodeID, "pid", os.Getpid())
	auth := strings.Split(em.etcdAuth, ":")
	if len(auth) != 2 {
		err := fmt.Errorf("etcd auth info must be in format user:pass")
		slog.Error(err.Error(), "value", em.etcdAuth)
		return err
	}

	endpoints := strings.Split(em.etcdAddr, ",")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
		Username:  auth[0],
		Password:  auth[1],
	})
	if err != nil {
		slog.Error("connect to etcd failed", "error", err)
		return err
	}
	defer cli.Close()

	session, err := concurrency.NewSession(cli)
	defer session.Close()

	election := concurrency.NewElection(session, em.electName)

	go em.ping(election)

	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		cmd := scaner.Text()
		switch cmd {
		case cmdElect:
			go em.elect(election)
		case cmdProclaim:
			go em.proclaim(election)
		case cmdResign:
			go em.resign(election)
		case cmdWatch:
			go em.watch(election)
		case cmdQuery:
			go em.query(election)
		case cmdRevision:
			go em.revision(election)
		default:
			slog.Error("unknown input command")
		}
	}
	return nil
}

func (em *ElectManager) value() string {
	return fmt.Sprintf("value-%d-%03d", em.nodeID, em.round)
}

func (em *ElectManager) elect(election *concurrency.Election) error {
	slog.Info("start electing", "node", em.nodeID)
	if err := election.Campaign(context.Background(), em.value()); err != nil {
		slog.Error("campaign failed", "error", err)
		return err
	}
	em.round++

	// header := election.Header()
	// slog.Info("current header info", "clusterID", header.ClusterId,
	// 	"memberID", header.MemberId, "revision", header.Revision)

	return nil
}

func (em *ElectManager) proclaim(election *concurrency.Election) error {
	slog.Info("start proclaiming", "node", em.nodeID)
	if err := election.Proclaim(context.Background(), em.value()); err != nil {
		slog.Error("proclaim failed", "error", err)
		return err
	}
	em.round++
	return nil
}

func (em *ElectManager) resign(election *concurrency.Election) error {
	slog.Info("start resigning", "node", em.nodeID)
	if err := election.Resign(context.Background()); err != nil {
		slog.Error("resign failed", "error", err)
		return err
	}
	return nil
}

func (em *ElectManager) watch(election *concurrency.Election) error {
	ch := election.Observe(context.Background())
	for {
		select {
		case res := <-ch:
			slog.Info("watching current leader info", "key", res.Kvs[0].Key, "value", res.Kvs[0].Value)
		}
	}
}

func (em *ElectManager) query(election *concurrency.Election) error {
	res, err := election.Leader(context.Background())
	if err != nil {
		slog.Error("query leader failed", "error", err)
		return err
	}
	slog.Info("current leader info", "key", res.Kvs[0].Key, "value", res.Kvs[0].Value)

	// header := election.Header()
	// slog.Info("current header info", "clusterID", header.ClusterId,
	// 	"memberID", header.MemberId, "revision", header.Revision)

	return nil
}

func (em *ElectManager) revision(election *concurrency.Election) error {
	rev := election.Rev()
	slog.Info("current leader", "revision", rev)
	return nil
}

func (em *ElectManager) ping(election *concurrency.Election) {
	tiker := time.NewTicker(time.Second * 30)
	defer tiker.Stop()
	for {
		select {
		case <-tiker.C:
			res, err := election.Leader(context.Background())
			if err != nil {
				slog.Error("query leader failed", "error", err)
				continue
			}
			value := string(res.Kvs[0].Value)
			entry := strings.Split(value, "-")
			isLeader := entry[1] == strconv.Itoa(em.nodeID)
			slog.Info("current leader info", "key", res.Kvs[0].Key, "value", value, "isLeader", isLeader)
		}
	}
}
