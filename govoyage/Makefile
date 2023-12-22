gen:
	@protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.
	@echo "\033[32mprotoc output files generated at: pbgen"
	@ls -lrt pbgen | grep pb.go


clean:
	@rm -rf ./pbgen/*


test:
	go test -cover -race .


server:
	go run cmd/server/server.go --port 8080


client:
	go run cmd/client/client.go --address 0.0.0.0:8080


.PHONY: gen clean server client test
