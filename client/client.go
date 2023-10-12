// client.go
package main

import (
	"context"
	"log"

	pb "github.com/atulhere/go-gRPC/helloworld"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Maa"})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	log.Printf("Response: %s", response.Message)
}
