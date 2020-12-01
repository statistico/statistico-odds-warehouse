package main

import (
	"github.com/statistico/statistico-odds-warehouse/internal/bootstrap"
	"github.com/statistico/statistico-odds-warehouse/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	app := bootstrap.BuildContainer(bootstrap.BuildConfig())

	opts := grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle:5*time.Minute})
	server := grpc.NewServer(opts)

	proto.RegisterMarketServiceServer(server, app.MarketService())

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

