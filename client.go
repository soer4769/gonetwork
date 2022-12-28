package main

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"time"

	"github.com/gonetwork/proto"
	"google.golang.org/grpc"
)

type ShakeType string

const (
	SYN ShakeType = "SYN"
	ACK ShakeType = "ACK"
)

// Representation of the 3-way handshake.
type Flags struct {
	SYN, ACK, FIN bool
}

// Repesentation of a TCP package.
type Pack struct {
	SeqNum, AckNum uint32
	Message        string
	Status         Flags
}

// Function for sending message utilizing a pack.
func SendMessage(c TCPHandshake.HandshakeClient, p Pack) (*TCPHandshake.TCPPack, error) {

	// Initiating the data to be sent corresponding to the pack.
	r, err := c.ConnSend(context.Background(), &TCPHandshake.TCPPack{
		SeqNum:  p.SeqNum,
		AckNum:  p.AckNum,
		Message: p.Message,
		Status: &TCPHandshake.Flags{
			SYN: p.Status.SYN,
			ACK: p.Status.ACK,
			FIN: p.Status.FIN,
		},
	})

	if err != nil {
		log.Fatalf("could not handshake: %v", err)
	}

	// Send/return the message.
	return r, err
}

func createPacket(msg ShakeType) Pack {

	if msg == SYN {
		return Pack{SeqNum: rand.Uint32(), Message: "SYN", Status: Flags{SYN: true}}
	}

	return Pack{SeqNum: rand.Uint32(), Message: "ACK", Status: Flags{ACK: true}}
}

// Perform handshake.
func Shake(c TCPHandshake.HandshakeClient) {
	log.Printf("Establishing Simulated TCP connection with server...")

	// Send a packet to the server with SYN set to 1 (true); SYN packet.
	syn := createPacket(SYN)
	log.Printf("Sending message to server:\n\t%+v\n", syn)

	r, err := SendMessage(c, syn)
	if err != nil {
		return
	}

	// If no errors occured upon sending the message, server received the message.
	log.Printf("Received message from server:\n\t%+v\n", r)

	// Server responds with SYN-ACK packet. Client sends back an ACK packet.
	ack := createPacket(ACK)
	log.Printf("sending message to server:\n\t%+v\n", ack)

	r, err = SendMessage(c, ack)
	if err != nil {
		return
	}

	log.Printf("Simulated TCP handshake successfully connected...")
}

func main() {

	// Changes the seed at start up so rand.Uint32 isn't the same every time.
	rand.Seed(time.Now().UnixNano())

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// New client.
	c := TCPHandshake.NewHandshakeClient(conn)

	// Commence a handshake.
	Shake(c)
	err = conn.Close()
	if err != nil {
		return
	}
}
