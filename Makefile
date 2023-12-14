exec_rest:
	go run ./cmd/rest

exec_grpc:
	go run ./cmd/grpc

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