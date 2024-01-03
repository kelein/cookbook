MODNAME="github.com/kelein/cookbook/govoyage"
Branch?=`git rev-parse --abbrev-ref HEAD`
SHA1=`git rev-parse --short HEAD`
Date=`date +%Y-%m-%dT%H:%M:%S`
Version=$(Branch)@$(SHA1)@$(Date)
User=`whoami`@`hostname`


LDFLAGS=-ldflags "-X '${MODNAME}/pkg/version.Version=${Version}'	\
	-X '${MODNAME}/pkg/version.Revision=${SHA1}'					\
	-X '${MODNAME}/pkg/version.Branch=${Branch}'					\
	-X '${MODNAME}/pkg/version.BuildDate=${Date}'					\
	-X '${MODNAME}/pkg/version.BuildUser=${User}'"


app:
	@echo "\033[32mcurrent build flag version: ${Version}\033[0m"
	@go build ${LDFLAGS} -o bin/govoyage-server ./cmd/server/server.go
	@echo "\033[32mbinary file output target at: bin/govoyage-server\033[0m"


app-linux:
	@echo "\033[32mcurrent build flag version: ${Version}\033[0m"
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/govoyage-server ./cmd/server/server.go
	@echo "\033[32mbinary file output target at: bin/govoyage-server\033[0m"


build: app


gen:
	@protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapi_out=naming=json=paths=source_relative:./assets
	@echo "\033[32mprotoc output files generated at: pbgen"
	@ls -lrt pbgen | grep pb.go
	@ls -lrt assets | grep yaml


clean:
	@rm -rf ./pbgen/*


test:
	go test -cover -race .


server:
	go run cmd/server/server.go --port 8080


client:
	go run cmd/client/client.go --address 0.0.0.0:8080


.PHONY: gen clean server client test


default: gen
