package models

import (
	"gopos/storage"
)

func Migrate() error {
	storage.ApplicationDB.AutoMigrate(&Product{})
	return nil
}
