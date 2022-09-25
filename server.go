package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "github.com/gonetwork/proto"
)

var (
    port = flag.Int("port", 50051, "The server port")
)

type server struct {
    TCPHandshake.UnimplementedHandshakeServer
}

func (s *server) ConnSend(ctx context.Context, in *TCPHandshake.SYN) (*TCPHandshake.ACK, error) {
    log.Printf("Received: %v", in.Num)
    return &TCPHandshake.ACK{Num: in.Num}, nil
}

func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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
