package endpoints

import (
	"gopos/repos"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	products := repos.GetAllProducts()
	return c.Status(http.StatusOK).JSON(products)
}
