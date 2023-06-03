package models

import (
	"gopos/storage"
)

func Migrate() error {
	storage.ApplicationDB.Db.AutoMigrate(&Product{})
	return nil
}
