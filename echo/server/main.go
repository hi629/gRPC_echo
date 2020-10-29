package main

import (
	"log"
	"net"

	pb "github.com/hi629/gRPC_echo/echo/proto"
	"google.golang.org/grpc"
)

func int() {
	log.SetFlags(0)
	log.SetPrefix("[echo]")
}

func main() {
	port := ":50051"
	// TCPのリスナを作成
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	// grpcサーバーを作成
	srv := grpc.NewServer()
	// gRPCサーバーに登録
	pb.RegisterEchoServiceServer(srv, &echoService{})
	log.Printf("start server on port %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
