package TCP_template

// responsible for defining TCP packages and their flags.

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
	FIN    ShakeType = "FIN"
)
