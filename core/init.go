package core

import (
	"gopos/models"
	"gopos/storage"

	"github.com/gofiber/fiber/v2"
)

func Init() {
	storage.InitializeDB()
	models.Migrate()
	app := fiber.New()
	SetupRoutes(app)

}
