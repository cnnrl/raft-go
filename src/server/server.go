package server

import (
	"context"
	"log"
	"net"

	"github.com/cnnrl/raft/src/node"
	pb "github.com/cnnrl/raft/src/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRaftServer
	node *node.Node
}

func (s *server) AppendEntries(_ context.Context, in *pb.Entry) (*pb.EntryResponse, error) {
	log.Printf("Received %v", in.GetMsg())
	s.node.Reset <- struct{}{}
	return &pb.EntryResponse{Recv: !s.node.Leading}, nil
}

func (s *server) RequestVote(_ context.Context, in *pb.VoteRequest) (*pb.Vote, error) {
	log.Printf("Received: %v", in.GetMsg())
	s.node.Reset <- struct{}{}
	return &pb.Vote{Vote: !s.node.Leading}, nil
}

func Start(n *node.Node, port string) {
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("could not listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRaftServer(s, &server{node: n})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
