build:
	protoc --proto_path=. --go_out=plugins=grpc:. internal/grpc/proto/*.proto