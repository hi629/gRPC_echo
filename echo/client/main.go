package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/hi629/gRPC_echo/echo/proto"
	"google.golang.org/grpc"
)

func main() {
	target := "localhost:50051"
	// サーバーの接続先を渡す。connにはgrpc.clientconnがはいる
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connnect: %s", err)
	}
	// 最後にconnectionをclose
	defer conn.Close()
	// grpc.ClientConnを渡す
	client := pb.NewEchoServiceClient(conn)
	msg := os.Args[1]
	ctx, cancel := context.WithTimeout(
		context.Background(), time.Second)
	defer cancel()
	r, err := client.Echo(ctx, &pb.EchoRequest{Message: msg})
	if err != nil {
		log.Println(err)
	}
	log.Println(r.GetMessage())
}
