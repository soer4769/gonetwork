package TCP_template

import (
	TCPHandshake "github.com/gonetwork/proto"
	"math/rand"
)

type Flags struct {
	SYN, ACK, FIN bool
}

type Pack struct {
	SeqNum, AckNum uint32
	Message        ShakeType
	Status         Flags
}

type ShakeType string

const (
	SYN    ShakeType = "SYN"
	ACK    ShakeType = "ACK"
	SYNACK ShakeType = "SYN+ACK"
)

func CreatePacket(msg ShakeType) Pack {

	if msg == SYN {
		return Pack{SeqNum: rand.Uint32(), Message: SYN, Status: Flags{SYN: true}}
	}

	return Pack{SeqNum: rand.Uint32(), Message: ACK, Status: Flags{ACK: true}}
}

func CreateSYNACKPacket(in *TCPHandshake.TCPPack) Pack {

	return Pack{SeqNum: rand.Uint32(), AckNum: in.SeqNum + 1, Message: SYNACK, Status: Flags{SYN: true, ACK: true}}
}
