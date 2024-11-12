package model

import "github.com/google/uuid"

type Guest struct {
	Ip string
}

func NewGuest(ip string) *Guest {
	return &Guest{Ip: ip}
}

func (g Guest) Register(credential *Credential) *User {
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
