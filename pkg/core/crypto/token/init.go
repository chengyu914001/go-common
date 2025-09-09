package token

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
	jwtMethod  = jwt.SigningMethodEdDSA
)

func init() {
	loadEd25519PrivateKey()
	loadEd25519PublicKey()
}

func loadEd25519PrivateKey() {
	pemData := os.Getenv("JWT_PRIVATE_KEY")
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		panic("invalid PEM block")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	privateKey = key.(ed25519.PrivateKey)
}

func loadEd25519PublicKey() {
	pemData := os.Getenv("JWT_PUBLIC_KEY")
	block, _ := pem.Decode([]byte(pemData))
	if block == nil {
		panic("invalid PEM block")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey = key.(ed25519.PublicKey)
}

func getPublicKey(token *jwt.Token) (any, error) {
	if token.Method != jwtMethod {
		return nil, ErrInvalidToken
	}
	return publicKey, nil
}
