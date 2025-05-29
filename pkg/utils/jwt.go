package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
	"strconv"
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

	priInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}


	privateKey, ok := priInterface.(*rsa.PrivateKey)
	if !ok {
		log.Fatal("not RSA private key")
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

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	publicKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		log.Fatal("not RSA public key")
	}

	return publicKey
}

func GenerateJWT(userID *uint, email string) (string, error) {
   var userIDStr string
    if userID != nil {
        userIDStr = strconv.FormatUint(uint64(*userID), 10)
    } else {
        userIDStr = ""
    }
    token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
 		"user_id": userIDStr,
        "email":   email,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString(GetPrivateKey())
}