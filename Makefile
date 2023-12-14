ENV := release
SRC_REST := ./cmd/rest/main.go
BIN_REST := ./bin/rest
BUILD_REST_CMD := go build -o $(BIN_REST) $(SRC_REST)

BIN_MIGRATE := ./migrate

migrate:
	go run ./cmd/migrate/migrate.go

exec_rest:
	go run ./cmd/rest

exec_grpc:
	go run ./cmd/grpc

migrate_docker:
	@$(BIN_MIGRATE)

protoc_grpc:
	protoc --go_out=. --go-grpc_out=. ./proto/*.proto

test:
	go test -v ./... -coverprofile cover.out
	go tool cover -html=cover.out	

build_docker:
	docker compose run --build

docker_refresh:
	docker compose down --volumes

docker_run:
	docker compose run