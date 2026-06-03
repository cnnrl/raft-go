package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cnnrl/raft/src/pb"
)

// func main() {
// 	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
//
// 	client := pb.NewGreeterClient(conn)
//
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Greeting: %s", r.GetMessage())
// }
