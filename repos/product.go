package repos

import (
	"fmt"
	"gopos/models"
	"gopos/storage"

	"github.com/oklog/ulid/v2"
)

func GetAllProducts() []models.Product {
	products := []models.Product{}
	storage.ApplicationDB.Find(&products)
	fmt.Println(products)
	return products
}
func CreateProduct(code string, price uint) models.Product {
	product := models.Product{
		Id:    ulid.Make(),
		Code:  code,
		Price: price,
	}
	storage.ApplicationDB.Create(&product)
	return product
}
