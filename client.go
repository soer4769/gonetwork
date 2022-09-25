package main

import (
    "context"
    "flag"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "github.com/gonetwork/proto"
)

var (
    addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
    flag.Parse()

    conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := TCPHandshake.NewHandshakeClient(conn)


    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.ConnSend(ctx, &TCPHandshake.SYN{Num: 0})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    log.Printf("Greeting: %s", r.Num)
}
