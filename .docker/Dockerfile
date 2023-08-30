# Step 1
FROM golang:1.19-alpine AS builder
RUN apk update && apk add git gcc g++
WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/console
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/grpc

# Step 2
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /opt
COPY ./bin ./bin
COPY ./database ./database
COPY ./.docker/wait-for-it.sh .
COPY --from=builder /go/bin/goose /usr/local/bin
COPY --from=builder /app/console .
COPY --from=builder /app/grpc .

CMD ["/bin/sh"]
