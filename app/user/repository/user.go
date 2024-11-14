package repository

import (
	"errors"
	"fmt"
	"gin-seed/app/database/connection"
	"gin-seed/app/user/entity"
	"gin-seed/app/user/model"
	"log"

	"gorm.io/gorm"
)

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
	db := connection.GetConnection()
	result := db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&entity.User{
			Id: user.Id,
		})
		tx.Create(&entity.Credential{
			UserId:   user.Credential.UserId,
			Username: user.Credential.Username,
			Password: user.Credential.Password,
		})
		return nil
	})

	if result != nil {
		log.Panic(result)
	}

	return nil
}

func GetByUsername(username string) *model.User {
	db := connection.GetConnection()

	var credential entity.Credential

	if err := db.Where(&entity.Credential{Username: username}).Take(&credential).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		log.Panic(err)
	}

	user := entity.User{Id: credential.UserId}
	if err := db.Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		log.Panic(err)
	}

	return &model.User{
		Guest: model.Guest{},
		Id:    user.Id,
		Credential: model.Credential{
			UserId:   user.Id,
			Username: credential.Username,
			Password: credential.Password,
		},
	}
}
