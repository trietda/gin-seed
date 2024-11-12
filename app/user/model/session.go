package model

import (
	"gin-seed/app/auth/service"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RefreshToken = [16]byte

type SessionMetadata struct {
	Ip string
}

type Session struct {
	Id           string
	UserId       string
	RefreshToken RefreshToken
	Metadata     *SessionMetadata
}

func NewSession(userId string, meta *SessionMetadata) *Session {
	return &Session{
		Id:           uuid.NewString(),
		UserId:       userId,
		RefreshToken: uuid.New(),
	}
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
