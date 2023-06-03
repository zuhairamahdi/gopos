package models

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id    ulid.ULID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	Code  string    `json:"code"`
	Price uint      `json:"price"`
}
