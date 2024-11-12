package model

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	jwt.RegisteredClaims
	SessionId string `json:"sessionid"`
}
