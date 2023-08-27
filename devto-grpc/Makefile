gen:
	@protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:.
	@echo "\033[32mprotoc output files generated at: repo\033[0m"
	@ls -lrt repo | grep pb.go


clean:
	@rm -rf ./repo/*


run:
	go run main.go


test:
	go test -cover -race ./convert
