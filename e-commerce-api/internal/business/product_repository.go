package business

import "github.com/jmaciel33/currency-exchange/e-commerce-api/internal/models"

type ProductRepository interface {

	Save(product models.Product) error

	Update(product models.Product) error

	FindById(value string) (models.Product, error)

	ListAllProducts() []models.Product

	Delete(value string) error

}