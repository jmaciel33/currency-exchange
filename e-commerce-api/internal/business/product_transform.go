package business

import (
	"github.com/google/uuid"
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/models"
)

func AddProductToEntity(data models.AddProduct) models.Product {
	return models.Product{
		Id: uuid.Must(uuid.NewRandom()),
		Name: data.Name,
		Price: data.Price,
	}
}

func UpdateProductToEntity(data models.Product) models.Product {
	return models.Product{
		Id: data.Id,
		Name: data.Name,
		Price: data.Price,
	}
}