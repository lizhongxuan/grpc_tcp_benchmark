package main

import (
	"context"
	"log"
	"google.golang.org/grpc"
	pb "grpc_tcp_test/proto/helloworld"
	"fmt"
	"grpc_tcp_test/tcp"
)

func main() {
	tcp_client()
}

var client *tcp.TcpConn


func init()  {
	client = tcp.NewClient("127.0.0.1:3333")
}

func tcp_client()  {
	client.Write([]byte("hello\n"))
}


const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)
var c pb.GreeterClient
var name = defaultName
var ctx context.Context
func grpc_init()  {
	fmt.Println("grpc client...")
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c = pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
}


func grpc_client() {

	_, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//log.Printf("Greeting: %s", r.GetMessage())
}