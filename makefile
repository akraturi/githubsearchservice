dev-setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.6
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.6

gen: dev-setup
	protoc --go_out=gen --go-grpc_out=gen api/v1/github-search.proto

golang-lint-run:
	golangci-lint run $(if $(TIMEOUT_MIN),--timeout=$(TIMEOUT_MIN)m)