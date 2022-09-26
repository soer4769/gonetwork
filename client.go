package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"github.com/gonetwork/proto"
	"google.golang.org/grpc"
)

func Shake(c TCPHandshake.HandshakeClient) {
	log.Printf("Establishing TCP connection with server...")

	for i := int32(0); i < 3; i++ {
		r, err := c.ConnSend(context.Background(), &TCPHandshake.SYN{Num: i})
		if err != nil {
			log.Fatalf("could not handshake: %v", err)
			return
		}
		log.Printf("handshake %d", r.Num)
	}

	log.Printf("TCP handshake successful.")
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := TCPHandshake.NewHandshakeClient(conn)

	Shake(c)
	err = conn.Close()
	if err != nil {
		return
	}
}
