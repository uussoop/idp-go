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
func ByteToRs512Priv(privBytes []byte) (*rsa.PrivateKey, error) {

	// Decode PEM encoded public key
	block, _ := pem.Decode(privBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	// Parse public key
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// Type assert to RSA public key
	privKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return privKey, nil

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
func Rs512PrivToByte(priv *rsa.PrivateKey) ([]byte, error) {

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, err
	}
	privPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: privBytes,
	})
	return privPEM, nil
}
func CompareRsaPublicKeys(pub1, pub2 *rsa.PublicKey) bool {
	if pub1.N.Cmp(pub2.N) == 0 && pub1.E == pub2.E {
		return true
	}
	return false
}
func CompareRsaPrivKeys(priv1, priv2 *rsa.PrivateKey) bool {
	if priv1.N.Cmp(priv2.N) == 0 && priv1.E == priv2.E {
		return true
	}
	return false
}
