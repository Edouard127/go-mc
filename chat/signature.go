package chat

import (
	"crypto/rand"
	"encoding/binary"
	pk "github.com/Edouard127/go-mc/net/packet"
	"io"
	"time"
)

type SignedMessage struct {
	Message      string
	Timestamp    int64
	Salt         int64
	HasSignature bool
	Signature    []byte
	Count        int
	Ack          pk.FixedBitSet
}

func SignMessage(msg Message, signature []byte) *SignedMessage {
	var salt int64
	binary.Read(rand.Reader, binary.BigEndian, &salt)

	return &SignedMessage{
		Message:      msg.Text,
		Timestamp:    time.Now().Unix(),
		Salt:         salt,
		HasSignature: false,
		Signature:    nil,
		Count:        0,
		Ack:          pk.NewFixedBitSet(20),
	}
}

func (s SignedMessage) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.String(s.Message),
		pk.Long(s.Timestamp),
		pk.Long(s.Salt),
		pk.Boolean(s.HasSignature),
		pk.Optional[pk.ByteArray, *pk.ByteArray]{
			Has:   s.Signature != nil,
			Value: s.Signature,
		},
		pk.VarInt(s.Count),
		pk.Optional[pk.FixedBitSet, *pk.FixedBitSet]{
			Has:   s.Count > 0,
			Value: s.Ack,
		},
	}.WriteTo(w)
}
