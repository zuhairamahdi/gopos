package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var ApplicationDB DbInstance

func InitializeDB() {
	DNS := "host=localhost user=postgres password=password dbname=gopos port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	db.Logger.LogMode(logger.Info)
	ApplicationDB = DbInstance{
		Db: db,
	}
}
