package core

import (
	"gopos/models"
	"gopos/storage"
)

func Init() {
	storage.InitializeDB()
	models.Migrate()

}
