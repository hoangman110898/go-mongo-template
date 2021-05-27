package services

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

var (
	privateKey *rsa.PrivateKey
)

func GenPrivateKey(){
	rng := rand.Reader
	var err error
	privateKey, err = rsa.GenerateKey(rng, 2048)
	if err != nil {
		log.Fatalf("rsa.GenerateKey: %v", err)
	}
}

func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}