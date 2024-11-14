package entity

import (
	"gin-seed/app/database/entity"
)

type User struct {
	entity.Base
  Id string `gorm:"primaryKey"`
}
