FROM golang:1.18.10-alpine

WORKDIR /app/rest

COPY go.mod go.sum ./
RUN rm -rf vendor/* bin/*

RUN go clean -mod=mod
RUN go mod tidy
RUN go mod download && go mod verify

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./cmd/migrate/
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/rest/

# FROM alpine:3 as migration

# COPY --from=rest /app/migrate /app/
# COPY --from=rest /app/Makefile /app/

# CMD make migrate_docker

# FROM alpine:3 as migration

# WORKDIR /app/rest

# COPY --from=builder /app/rest /app/

EXPOSE 8080

CMD ./main