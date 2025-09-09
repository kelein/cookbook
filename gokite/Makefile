REGISTRY="github.com/kelein/cookbook"
MODNAME="github.com/kelein/cookbook"
BRANCH?=`git rev-parse --abbrev-ref HEAD`
SHA1=`git rev-parse --short HEAD`
DATE=`date +%Y-%m-%dT%H:%M:%S`
VERSION=`git describe --tags --always`
USER=`whoami`@`hostname`


LDFLAGS=-ldflags "-X '${MODNAME}/pkg/version.AppVersion=${VERSION}'	\
	-X '${MODNAME}/pkg/version.Revision=${SHA1}'					\
	-X '${MODNAME}/pkg/version.Branch=${BRANCH}'					\
	-X '${MODNAME}/pkg/version.BuildDate=${DATE}'					\
	-X '${MODNAME}/pkg/version.BuildUser=${USER}'"


app:
	@echo "\033[32mcurrent build flag version: ${VERSION}\033[0m"
	go build ${LDFLAGS} -o bin/gokite cmd/gokite/gokite.go
	@echo "\033[32mbinary file output target at: bin/gokite\033[0m"
	@bin/gokite -v


app-linux:
	@echo "\033[32mcurrent build flag version: ${VERSION}\033[0m"
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/gokite cmd/gokite/gokite.go
	@echo "\033[32mbinary file output target at: bin/gokite\033[0m"
	@bin/gokite -v


.PHONY: generate
# generate wire code
generate:
# 	wire internal/inject/wire.go
	@/bin/rm -rf kitgen
	go generate ./...
	kitex -I proto -gen-path kitgen -module ${MODNAME} proto/*.proto


.PHONY: init
# init build packages
init:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/google/wire/cmd/wire@latest
	go mod tidy


.PHONY: swag
# build server binary
build: app-linux


.PHONY: swag
# generate swag docs
swag:
	swag init -g cmd/gokite/gokite.go -o docs


.PHONY: image
# build docker image
image:
	docker build --platform=linux/amd64 -t ${REGISTRY}:${VERSION} -f Dockerfile .
	docker push ${REGISTRY}:${VERSION}


# show this make help
help:
	@echo 'Usage:'
	@echo '  make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m  %-16s\033[0m%s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
