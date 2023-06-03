package main

import (
	"gopos/storage"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id    ulid.ULID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	Code  string    `json:"code"`
	Price uint      `json:"price"`
}

func main() {

	storage.InitializeDB()

	storage.ApplicationDB.Db.AutoMigrate(&Product{})
	storage.ApplicationDB.Db.Create(&Product{Code: "SHA1", Price: 100, Id: ulid.Make()})
	//db.Create(&Product{Code: "D42", Price: 100})

}
