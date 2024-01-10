module github.com/kelein/cookbook

go 1.21

require (
	github.com/brianvoe/gofakeit/v6 v6.26.4
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.5.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0
	github.com/jedib0t/go-pretty/v6 v6.5.2
	github.com/jinzhu/copier v0.4.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.18.0
	github.com/stretchr/testify v1.8.4
	github.com/swaggest/swgui v1.7.5
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go.etcd.io/etcd/client/v3 v3.5.11
	go.etcd.io/etcd/server/v3 v3.5.11
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/zap v1.26.0
	golang.org/x/crypto v0.18.0
	golang.org/x/net v0.20.0
	golang.org/x/sync v0.6.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
	sourcegraph.com/sourcegraph/appdash v0.0.0-20211028080628-e2786a622600
)

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/basictracer-go v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/shurcooL/httpfs v0.0.0-20230704072500-f1e31cf0ba5c // indirect
	github.com/shurcooL/vfsgen v0.0.0-20230704071429-0000e147ea92 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/vearutop/statigz v1.4.0 // indirect
	github.com/xiang90/probing v0.0.0-20221125231312-a49e3df8f510 // indirect
	go.etcd.io/bbolt v1.3.8 // indirect
	go.etcd.io/etcd/api/v3 v3.5.11 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.11 // indirect
	go.etcd.io/etcd/client/v2 v2.305.11 // indirect
	go.etcd.io/etcd/pkg/v3 v3.5.11 // indirect
	go.etcd.io/etcd/raft/v3 v3.5.11 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.1 // indirect
	go.opentelemetry.io/otel v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/goleak v1.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sourcegraph.com/sourcegraph/appdash-data v0.0.0-20151005221446-73f23eafcf67 // indirect
)

replace (
	github.com/coreos/bbolt v1.3.8 => go.etcd.io/bbolt v1.3.8
	go.etcd.io/bbolt v1.3.8 => github.com/coreos/bbolt v1.3.8
	sourcegraph.com/sourcegraph/appdash-data v0.0.0-20151005221446-73f23eafcf67 => github.com/sourcegraph/appdash-data v0.0.0-20151005221446-73f23eafcf67
)
