package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc_tcp_test/proto/helloworld"
	"fmt"
	"grpc_tcp_test/tcp"
)

const (
	port = ":50051"
)


func main()  {
	tcp_server()
}

func tcp_server()  {
	server := tcp.NewServer(3333)
	for   {
		msg:= server.Read()
		fmt.Println("server read:",string(msg))
	}

}


// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func grpc_server() {
	fmt.Println("grpc server...")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

