FROM golang:1.18.10-alpine as builder

WORKDIR /app/grpc

COPY go.mod go.sum ./
RUN rm -rf vendor/* bin/*

RUN go clean -mod=mod
RUN go mod tidy
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./cmd/migrate/
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/grpc/

FROM golang:1.18.10-alpine as migration

COPY --from=grpc /app/migrate /app/
COPY --from=grpc /app/Makefile /app/

CMD make migrate_docker

FROM alpine:3 as migration

WORKDIR /app/grpc

COPY --from=builder /app/grpc /app/

EXPOSE 50053

CMD ./main
