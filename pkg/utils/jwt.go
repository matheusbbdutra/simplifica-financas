package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GetPrivateKey() *rsa.PrivateKey {
	keyData, err := os.ReadFile("config/jwt/private.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(keyData)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	return privateKey
}

func GetPublicKey() *rsa.PublicKey {
	keyData, err := os.ReadFile("config/jwt/public.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(keyData)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the key")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	return publicKey
}

func GenerateJWT(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    return token.SignedString(GetPrivateKey())
}