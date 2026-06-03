package main

import (
	"github.com/cnnrl/raft/src/client"
	"github.com/cnnrl/raft/src/node"
	pb "github.com/cnnrl/raft/src/pb"
	"github.com/cnnrl/raft/src/server"
	"os"
	"time"
)

func main() {

	port := os.Args[1]
	peers := os.Args[2:]
	n := node.Node{
		Port:    port,
		Conns:   []pb.RaftClient{},
		Reset:   make(chan struct{}, 1),
		Stop:    make(chan struct{}, 1),
		Leading: false,
	}

	go server.Start(&n, port)

	wait := time.NewTimer(time.Duration(10) * time.Second)
	select {
	case <-wait.C:
	}

	client.Connect(&n, peers)
	go client.StartHeart(&n)

	select {}
}
