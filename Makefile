.PHONY: clean protoc test grpc build

clean:
	@echo "--- Cleanup all build and generated files ---"
	@rm -vf pkg/infrastructure/grpc/proto/validator/*.pb.go
	@rm -vf pkg/infrastructure/grpc/proto/class/*.pb.go

protoc: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p pkg/infrastructure/grpc/proto/validator
	@mkdir -p pkg/infrastructure/grpc/proto/class

	@echo "--- Compiling all proto files ---"
	@cd ./pkg/shared/proto/validator && protoc -I. --go_out=../../../infrastructure/grpc/proto/validator --go-grpc_out=require_unimplemented_servers=false:../../../infrastructure/grpc/proto/validator *.proto
	@cd ./pkg/shared/proto/class && protoc -I. -I../validator --go_out=../../../infrastructure/grpc/proto/class --go-grpc_out=require_unimplemented_servers=false:../../../infrastructure/grpc/proto/class --govalidators_out=../../../infrastructure/grpc/proto/class *.proto

grpc:
	@echo "--- running gRPC server in dev mode ---"
	@go run cmd/server/grpc/main.go

setup:
	@echo " --- Setup and generate configuration --- "
	@cp internal/config/example/database.yml internal/config/db/database.yml
	@cp internal/config/example/server.yml internal/config/server/server.yml

build: setup protoc
	@echo "--- Building binary file ---"
	@go build -o ./main cmd/server/grpc/main.go