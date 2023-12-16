package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "grpc-helloworld/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "User"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to be used as parameter")
)

func logReply(request string, message string) {
	log.Printf("Reply for '%s', message: %s", request, message)
}

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	logReply("SayHello", r.GetMessage())

	r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet again: %v", err)
	}
	logReply("SayHelloAgain", r.GetMessage())

	rr, err := c.SayBye(ctx, &pb.ByeRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not say bye: %v", err)
	}
	logReply("SayBye", rr.GetMessage())
}
