package data

import (
	"encoding/base64"
	"encoding/pem"
	"fmt"
	pk "github.com/Edouard127/go-mc/net/packet"
	"io"
	"time"
)

type KeyPairResp struct {
	KeyPair struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	} `json:"keyPair"`
	PublicKeySignature   string    `json:"publicKeySignature"`
	PublicKeySignatureV2 string    `json:"publicKeySignatureV2"`
	ExpiresAt            time.Time `json:"expiresAt"`
	RefreshedAfter       time.Time `json:"refreshedAfter"`
}

func (k KeyPairResp) WriteTo(w io.Writer) (int64, error) {
	block, _ := pem.Decode([]byte(k.KeyPair.PublicKey))
	if block == nil {
		return 0, fmt.Errorf("pem decode error: no data is found")
	}
	signature, err := base64.StdEncoding.DecodeString(k.PublicKeySignatureV2)
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		pk.Long(k.ExpiresAt.UnixMilli()),
		pk.ByteArray(block.Bytes),
		pk.ByteArray(signature),
	}.WriteTo(w)
}
