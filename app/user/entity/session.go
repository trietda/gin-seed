package entity

import (
	"gin-seed/app/database/entity"
	"gin-seed/app/user/model"
)

type Session struct {
	entity.Base
	Id           string
	UserId       string
	RefreshToken string
	Metadata     *model.SessionMetadata `gorm:"serializer:json"`
}
