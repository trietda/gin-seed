package service

import (
	"crypto/ed25519"
	"log"
	"os"
	"sync"

	"github.com/golang-jwt/jwt/v5"
)

var (
	privateKeyLock = &sync.Mutex{}
	privateKey     ed25519.PrivateKey
	publicKeyLock  = &sync.Mutex{}
	publicKey      ed25519.PublicKey
)

func getPrivateKey() ed25519.PrivateKey {
	if privateKey == nil {
		privateKeyLock.Lock()
		defer privateKeyLock.Unlock()

		key, err := jwt.ParseEdPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))

		if err != nil {
			log.Fatalln("Invalid private key")
		}

		privateKey = key.(ed25519.PrivateKey)
	}

	return privateKey
}

func getPublicKey() ed25519.PublicKey {
	if publicKey == nil {
		publicKeyLock.Lock()
		defer publicKeyLock.Unlock()

		privateKey := getPrivateKey()
		publicKey = privateKey.Public().(ed25519.PublicKey)
	}

	return publicKey
}
