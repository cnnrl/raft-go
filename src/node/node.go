package node

import (
	pb "github.com/cnnrl/raft/src/pb"
)

type Node struct {
	Port    string
	Conns   []pb.RaftClient
	Reset   chan struct{}
	Stop    chan struct{}
	Leading bool
}
