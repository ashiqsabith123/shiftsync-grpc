form:
	protoc --go_out=. --go-grpc_out=. --proto_path=./pkg/form/pb/ form.proto

auth:
	protoc --go_out=. --go-grpc_out=. --proto_path=./pkg/auth/pb/ auth.proto

run:
	go run cmd/api/main.go