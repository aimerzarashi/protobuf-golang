package main

import (
	"context"
	"log"
	"net"
	"protobuf/internal/infrastructure/grpc/helloworld"
	"protobuf/internal/infrastructure/grpc/helloworld2"

	"google.golang.org/grpc"
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
	helloworld.RegisterGreeterServer(grpcServer, &hello{})
	helloworld2.RegisterGreeterServer(grpcServer, &hello2{})

	// 以下でリッスンし続ける
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Print("main end")
}

type hello struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *hello) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received hello: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

type hello2 struct {
	helloworld2.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *hello2) SayHello(ctx context.Context, in *helloworld2.HelloRequest) (*helloworld2.HelloReply, error) {
	log.Printf("Received hello2: %v", in.GetName())
	return &helloworld2.HelloReply{Message: "Hello " + in.GetName()}, nil
}