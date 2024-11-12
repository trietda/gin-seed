package repository

import (
	"fmt"
	"gin-seed/app/user/model"
)

var users map[string]model.User = make(map[string]model.User)
var credentials map[string]*model.Credential = make(map[string]*model.Credential)

const (
	UserExistedError = iota
)

type SaveUserError struct {
	Code int
}

func (e SaveUserError) Error() string {
	return fmt.Sprintf("%d", e.Code)
}

func SaveUser(user model.User) *SaveUserError {
	if _, ok := credentials[user.Credential.Username]; ok {
		return &SaveUserError{Code: UserExistedError}
	}

	users[user.Id] = user
	credentials[user.Credential.Username] = &user.Credential
	return nil
}

func GetByUsername(username string) *model.User {
	credential, ok := credentials[username]

	if !ok {
		return nil
	}

	user := users[credential.UserId]
	return &user
}
