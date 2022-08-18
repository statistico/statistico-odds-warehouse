# Step 1
FROM golang:1.17-alpine AS builder
RUN apk update && apk add git gcc g++
WORKDIR /app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/grpc
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/queue

# Step 2
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /opt
COPY ./bin ./bin
COPY ./database ./database
COPY ./.docker/wait-for-it.sh .
COPY --from=builder /go/bin/goose /usr/local/bin
COPY --from=builder /app/grpc .
COPY --from=builder /app/queue .

CMD ["/bin/sh"]
