package client

import (
	"context"
	"log"
	"math/rand/v2"
	"time"

	"github.com/cnnrl/raft/src/node"
	pb "github.com/cnnrl/raft/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(n *node.Node, ports []string) {
	for _, port := range ports {
		conn, err := grpc.NewClient("localhost:"+port,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("could not connect to peer %s: %v", port, err)
			continue
		}
		n.Conns = append(n.Conns, pb.NewRaftClient(conn))
	}
}

func StartHeart(n *node.Node) {
	timer := time.NewTimer(randTime())
	for {
		select {
		case <-timer.C:
			if n.Leading {
				appendEntries(n)
				timer.Reset(30 * time.Millisecond)
			} else {
				if requestVote(n) {
					n.Leading = true
					timer.Reset(30 * time.Millisecond)
				} else {
					timer.Reset(randTime())
				}
			}
		case <-n.Reset:
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(randTime())
		case <-n.Stop:
			return
		}
	}
}

func requestVote(n *node.Node) bool {
	total := 0
	for _, conn := range n.Conns {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		r, err := conn.RequestVote(ctx, &pb.VoteRequest{Msg: "pls vote"})

		cancel()

		if err != nil {
			log.Printf("error requesting vote: %v", err)
			continue
		}
		if r.GetVote() {
			total++
		}
	}
	return total >= len(n.Conns)/2
}

func randTime() time.Duration {
	return time.Duration(rand.N(150)+150) * time.Millisecond
}

func appendEntries(n *node.Node) {
	for _, conn := range n.Conns {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		r, err := conn.AppendEntries(ctx, &pb.Entry{Msg: "log/heartbeat"})

		cancel()

		if err != nil {
			log.Printf("error appending entries: %v", err)
		}
		if !r.GetRecv() {
			//TODO: think of way to handle not getting a vote
		}
	}
}
