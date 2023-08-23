package data

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	pk "github.com/Edouard127/go-mc/net/packet"
	"io"
	"strconv"
	"time"
)

type KeyPair struct {
	Pair struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	} `json:"keyPair"`
	PublicKeySignature   string    `json:"publicKeySignature"`
	PublicKeySignatureV2 string    `json:"publicKeySignatureV2"`
	ExpiresAt            time.Time `json:"expiresAt"`
	RefreshedAfter       time.Time `json:"refreshedAfter"`
	UUID                 string    `json:"-"` // This is not returned by the API but required for the client
}

func (k KeyPair) WriteTo(w io.Writer) (int64, error) {
	block, _ := pem.Decode([]byte(k.Pair.PublicKey))
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

func (k KeyPair) ToSession(uuid string) PlayerSession {
	return PlayerSession{
		KeyPair: k,
		UUID:    uuid,
	}
}

type PlayerSession struct {
	KeyPair
	UUID string
}

func (k PlayerSession) WriteTo(w io.Writer) (int64, error) {
	pubBlock, _ := pem.Decode([]byte(k.Pair.PublicKey))
	if pubBlock == nil {
		panic("failed to parse PEM pubBlock containing the public key")
	}

	hash := sha1.New()
	hash.Write([]byte(k.UUID))
	hash.Write([]byte(strconv.FormatInt(k.ExpiresAt.UnixMilli(), 10)))
	hash.Write(pubBlock.Bytes)

	privBlock, _ := pem.Decode([]byte(k.Pair.PrivateKey))
	der, err := x509.ParsePKCS8PrivateKey(privBlock.Bytes)
	if err != nil {
		panic("failed to parse DER encoded private key: " + err.Error())
	}

	signature, err := rsa.SignPKCS1v15(nil, der.(*rsa.PrivateKey), crypto.SHA1, hash.Sum(nil))
	if err != nil {
		panic("failed to sign hash: " + err.Error())
	}

	return pk.Tuple{
		pk.Long(k.ExpiresAt.UnixMilli()),
		pk.ByteArray(pubBlock.Bytes),
		pk.ByteArray(signature),
	}.WriteTo(w)
}
