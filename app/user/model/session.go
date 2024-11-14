package model

import (
	"gin-seed/app/auth/service"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IJwtService interface {
	NewJwt(jwt.Claims) string
}

type SessionMetadata struct {
	Ip string `json:"ip"`
}

type Session struct {
	Id           string
	UserId       string
	RefreshToken string
	Metadata     *SessionMetadata
}

func (s *Session) GenerateAccessToken() (string, error) {
	claim := UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-seed",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		SessionId: s.Id,
	}
	return service.NewJwt(claim)
}
