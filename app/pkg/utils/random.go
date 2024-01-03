package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return b
}

func GeneratePairKey() (private *rsa.PrivateKey) {
	private, _ = rsa.GenerateKey(rand.Reader, 4096)

	return
}
func ByteToRs512Pub(pubBytes []byte) (*rsa.PublicKey, error) {

	// Decode PEM encoded public key
	block, _ := pem.Decode(pubBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	// Parse public key
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Type assert to RSA public key
	pubKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return pubKey, nil

}
func Rs512PubToByte(pub *rsa.PublicKey) ([]byte, error) {

	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}
	pubPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	})
	return pubPEM, nil
}
func CompareRsaPublicKeys(pub1, pub2 *rsa.PublicKey) bool {
	if pub1.N.Cmp(pub2.N) == 0 && pub1.E == pub2.E {
		return true
	}
	return false
}
