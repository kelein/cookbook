gen:
	@# @protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.

	@# @protoc --proto_path=proto proto/*.proto --go-grpc_out=./repo

	@protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

	@echo "\033[32mprotoc output files generated at: repo\033[0m"
	@ls -lrt repo | grep pb.go


clean:
	@rm -rf ./repo/*


run:
	go run main.go


test:
	go test -cover -race ./convert


server:
	go run cmd/server/main.go --port 8000


client:
	go run cmd/client/main.go --address 0.0.0.0:8000
