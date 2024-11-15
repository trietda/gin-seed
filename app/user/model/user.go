package model

import (
	"errors"

	"github.com/google/uuid"
)

type Guest struct {
	Ip string
}

func NewGuest(ip string) *Guest {
	return &Guest{Ip: ip}
}

func (g Guest) Register(credential *Credential) *User {
	if credential == nil {
    panic(errors.New("Empty credential"))
	}

	userId := uuid.NewString()
	credential.UserId = userId
	return &User{
		Guest:      g,
		Id:         userId,
		Credential: *credential,
	}
}

type User struct {
	Guest
	Id         string
	Credential Credential
}

func (u User) Login() *Session {
	return &Session{
		Id:           uuid.NewString(),
		UserId:       u.Id,
		RefreshToken: uuid.NewString(),
		Metadata: &SessionMetadata{
			Ip: u.Ip,
		},
	}
}
