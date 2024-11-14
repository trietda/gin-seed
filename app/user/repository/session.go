package repository

import (
	"errors"
	"gin-seed/app/database/connection"
	"gin-seed/app/user/entity"
	"gin-seed/app/user/model"
	"log"

	"gorm.io/gorm"
)

func SaveSession(session model.Session) {
	db := connection.GetConnection()

	result := db.Create(&entity.Session{
		Id:           session.Id,
		UserId:       session.UserId,
		RefreshToken: session.RefreshToken,
		Metadata:     session.Metadata,
	})

	if result.Error != nil {
		log.Panic(result.Error)
	}
}

func GetSession(refreshToken string) *model.Session {
	db := connection.GetConnection()

	var session entity.Session
	if err := db.Where(&entity.Session{RefreshToken: refreshToken}).Take(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		log.Panic(err)
	}

	return &model.Session{
		Id:           session.Id,
		UserId:       session.UserId,
		RefreshToken: session.RefreshToken,
		Metadata:     session.Metadata,
	}
}
