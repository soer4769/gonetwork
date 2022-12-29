package TCP_template

import (
	TCPHandshake "github.com/gonetwork/proto"
	"math/rand"
)

// responsible for creating TCP packets.

func CreateSYNPacket() Pack {
	return Pack{SeqNum: rand.Uint32(), Message: ACK, Status: Flags{ACK: true}}
}

func CreateACKPacket() Pack {
	return Pack{SeqNum: rand.Uint32(), Message: SYN, Status: Flags{SYN: true}}
}

func CreateSYNACKPacket(in *TCPHandshake.TCPPack) Pack {

	return Pack{SeqNum: rand.Uint32(), AckNum: in.SeqNum + 1, Message: SYNACK, Status: Flags{SYN: true, ACK: true}}
}
