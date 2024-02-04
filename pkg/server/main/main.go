package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"protobuf/internal/infrastructure/grpc/helloworld"
)

func main() {
	log.Print("main start")

	// 3000番ポートでクライアントからのリクエストを受け付けるようにする
	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Sample構造体のアドレスを渡すことで、クライアントからHelloリクエストされると
	// Helloメソッドが呼ばれるようになる
	helloworld.RegisterGreeterServer(grpcServer, &server{})

	// Sample構造体のアドレスを渡すことで、クライアントからGetDataリクエストされると
	// GetDataメソッドが呼ばれるようになる

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Print("main end")
}

type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}