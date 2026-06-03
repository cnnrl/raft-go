package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cnnrl/raft/src/pb"
	"github.com/cnnrl/raft/src/server"
)

func main() {
	port := os.Args[0]
	go server.Start(&port)
}
