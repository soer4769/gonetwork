package main

import (
	"context"
	packet "github.com/gonetwork/TCP_template"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/gonetwork/proto"
	"google.golang.org/grpc"
)

type server struct {
	TCPHandshake.UnimplementedHandshakeServer
}

// handle connection handshake for a client.
func (s *server) ConnSend(_ context.Context, in *TCPHandshake.TCPPack) (*TCPHandshake.TCPPack, error) {

	// If status is SYN, a client is trying to establish a connection.
	if in.Status.SYN {
		log.Printf("New client trying to establish simulated TCP connection...")
		log.Printf("Recieved message from client:\n\t%+v\n", in)

		// Respond with a SYN-ACK to the client.
		ack := packet.CreateSYNACKPacket(in)
		log.Printf("Sending response to client:\n\t%+v\n", ack)

		// send the SYN-ACK.
		return &TCPHandshake.TCPPack{
			SeqNum:  ack.SeqNum,
			AckNum:  ack.AckNum,
			Message: string(ack.Message),
			Status: &TCPHandshake.Flags{
				SYN: ack.Status.SYN,
				ACK: ack.Status.ACK,
			},
		}, nil
	}

	// If client responds with ACK, connection successfully established.
	if in.Status.ACK {
		log.Printf("Established simulated TCP connection succefully with client...")
	}

	return &TCPHandshake.TCPPack{}, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initiate new server
	s := grpc.NewServer()

	// Register the server as a handshake server
	TCPHandshake.RegisterHandshakeServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
