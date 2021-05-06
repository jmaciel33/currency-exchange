package factory

import (
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/app/config"
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/app/persistence"
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/business"
	_ "github.com/lib/pq"
)

var repository business.ProductRepository

func init() {
	db, err := config.InitDbConnection()
	if err != nil {
		panic(err)
	}
	repository = persistence.NewProductRepository(db)
}

func GetProductService() business.ProductService {
	return business.NewProductService(repository)
}
