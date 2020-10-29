package main

// importの書き方 https://qiita.com/taji-taji/items/5a4f17bcf5b819954cc1
import (
	//
	"context"
	// pbとしてエイリアスを作成。相対パスでも指定できるが推奨されていないためエラーになる
	// https://qiita.com/lp-peg/items/9f7da97d9f7292f16fbc
	pb "github.com/hi629/gRPC_echo/echo/proto"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.GetMessage()}, nil
}
