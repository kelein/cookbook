package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Cluster holds metadata + clientset
type Cluster struct {
	Name       string
	Kubeconfig string
	Client     *kubernetes.Clientset
}

// MultiClusterManager keeps track of multiple clusters
type MultiClusterManager struct {
	mu       sync.RWMutex
	clusters map[string]*Cluster
}

func NewMultiClusterManager() *MultiClusterManager {
	return &MultiClusterManager{
		clusters: make(map[string]*Cluster),
	}
}

// Register a cluster by kubeconfig file
func (m *MultiClusterManager) Register(name, kubeconfigPath string) error {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return fmt.Errorf("failed to load kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}

	cluster := &Cluster{
		Name:       name,
		Kubeconfig: kubeconfigPath,
		Client:     clientset,
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	m.clusters[name] = cluster
	return nil
}

// Get cluster info (version + nodes)
func (m *MultiClusterManager) Info(name string) error {
	m.mu.RLock()
	cluster, ok := m.clusters[name]
	m.mu.RUnlock()
	if !ok {
		return fmt.Errorf("cluster %s not found", name)
	}

	ctx := context.Background()
	ver, err := cluster.Client.Discovery().ServerVersion()
	if err != nil {
		return err
	}
	fmt.Printf("Cluster %s: Kubernetes %s\n", name, ver.String())

	nodes, err := cluster.Client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Cluster %s has %d nodes\n", name, len(nodes.Items))
	for _, n := range nodes.Items {
		fmt.Printf(" - %s (labels: %v)\n", n.Name, n.Labels)
	}

	comps, err := cluster.Client.CoreV1().ComponentStatuses().List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, comp := range comps.Items {
		fmt.Printf("Component %s: %s %s\n", comp.Name, comp.APIVersion, comp.String())
	}

	group, resources, err := cluster.Client.ServerGroupsAndResources()
	slog.Info("ServerGroupsAndResources", "group", group, "resources", resources)

	restr, err := json.Marshal(resources)
	if err != nil {
		return err
	}
	slog.Info("ServerGroupsAndResources Marshaled", "value", restr)

	info, err := cluster.Client.Discovery().ServerVersion()
	slog.Info("ServerVersion", "info", info)

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: multicluster <name> <kubeconfig-path>")
		return
	}
	name := os.Args[1]
	kubeconfig := os.Args[2]

	manager := NewMultiClusterManager()
	if err := manager.Register(name, kubeconfig); err != nil {
		panic(err)
	}

	if err := manager.Info(name); err != nil {
		panic(err)
	}
}
