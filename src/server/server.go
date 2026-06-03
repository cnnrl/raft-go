package server

import (
	"context"
	"log"
	"net"

	pb "github.com/cnnrl/raft/src/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRaftServer
}

func (s *server) AppendEntires(_ context.Context, in *pb.Entry) (*pb.EntryResponse, error) {
	log.Printf("Received %v", in.GetMsg())
	return &pb.EntryResponse{Recv: true}, nil
}

func (s *server) RequestVote(_ context.Context, in *pb.VoteRequest) (*pb.Vote, error) {
	log.Printf("Received: %v", in.GetMsg())
	return &pb.Vote{Vote: true}, nil
}

func Start(port *string) {
	lis, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
