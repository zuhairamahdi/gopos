package models

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       ulid.ULID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string
	Salt     string
}
