package service

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func NewJwt(claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claim)
	privateKey := getPrivateKey()
	return token.SignedString(privateKey)
}

func VerifyAccessToken(tokenString string, claim jwt.Claims) (*jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		claim,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodEd25519); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			publicKey := getPublicKey()
			return publicKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &token.Claims, nil
}
