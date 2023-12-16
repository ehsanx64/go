package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "grpc-helloworld/helloworld"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
}

func logRequest(request string, message string) {
	log.Printf("Request '%s', parameter: %v", request, message)
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logRequest("SayHello", in.GetName())
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s!!!", in.GetName())}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logRequest("SayHelloAgain", in.GetName())
	return &pb.HelloReply{Message: fmt.Sprintf("Hello again %s!", in.GetName())}, nil
}

func (s *server) SayBye(ctx context.Context, in *pb.ByeRequest) (*pb.ByeReply, error) {
	logRequest("SayBye", in.GetName())
	return &pb.ByeReply{Message: fmt.Sprintf("Bye %s!!!", in.GetName())}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
