package data

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	pk "github.com/Edouard127/go-mc/net/packet"
	"io"
	"time"
)

type PublicKey struct {
	ExpiresAt int64
	PublicKey *rsa.PublicKey
	Signature []byte
}

func (pub PublicKey) WriteTo(w io.Writer) (int64, error) {
	encoded, err := x509.MarshalPKIXPublicKey(pub.PublicKey)
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		pk.Long(pub.ExpiresAt),
		pk.ByteArray(encoded),
		pk.ByteArray(pub.Signature),
	}.WriteTo(w)
}

func (pub *PublicKey) ReadFrom(r io.Reader) (int64, error) {
	var (
		encoded   pk.ByteArray
		signature pk.ByteArray
	)
	n, err := pk.Tuple{
		&pub.ExpiresAt,
		&encoded,
		&signature,
	}.ReadFrom(r)
	if err != nil {
		return n, err
	}
	pub.PublicKey = decodedPublicKey(encoded)
	if err != nil {
		return n, err
	}
	pub.Signature = signature
	return n, nil
}

func (pub *PublicKey) Verify() bool {
	return VerifySignature(encodedPublicKey(pub.PublicKey), pub.Signature) && pub.ExpiresAt > time.Now().UnixMilli()
}

func VerifyMessage(hash, signature []byte) bool {
	return rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash, signature) == nil
}

func encodedPublicKey(pub *rsa.PublicKey) []byte {
	encoded, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		panic(err)
	}
	return encoded
}

func decodedPublicKey(encoded []byte) *rsa.PublicKey {
	decoded, err := x509.ParsePKIXPublicKey(encoded)
	if err != nil {
		panic(err)
	}
	return decoded.(*rsa.PublicKey)
}
