dev-setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen: dev-setup
	protoc --go_out=. --go-grpc_out=. api/v1/github-search.proto