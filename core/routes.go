package core

import (
	product "gopos/endpoints/product"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	setupProductsRoute(app)
	app.Listen(":3000")
}

func setupProductsRoute(app *fiber.App) {
	app.Get("/api/product/", product.GetAllProducts)
}
