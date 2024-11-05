server:
	go run cmd/main.go

proto.user:
	protoc -I=./pkg/common/proto --go_out=./pkg/user/pb/ --go-grpc_out=. pkg/common/proto/user-query.proto
	protoc -I=./pkg/common/proto --go_out=./pkg/user/pb/ --go-grpc_out=. pkg/common/proto/user-command.proto
proto:
	make proto.user
