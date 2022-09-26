package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "github.com/gonetwork/proto"
)

type server struct {
    TCPHandshake.UnimplementedHandshakeServer
}

func (s *server) ConnSend(ctx context.Context, in *TCPHandshake.SYN) (*TCPHandshake.ACK, error) {
    log.Printf("Received: %v", in.Num)
    return &TCPHandshake.ACK{Num: in.Num}, nil
}

func main() {
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()

    TCPHandshake.RegisterHandshakeServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
