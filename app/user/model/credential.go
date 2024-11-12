package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Credential struct {
	UserId   string
	Username string
	Password []byte
}

func NewCredential(username, password string) (*Credential, error) {
	hash, hashErr := bcrypt.GenerateFromPassword([]byte(password), 10)

	if hashErr != nil {
		return nil, hashErr
	}

	return &Credential{
		Username: username,
		Password: hash,
	}, nil
}

func (c Credential) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(c.Password, []byte(password))

	if err != nil {
		return false
	}

	return true
}
