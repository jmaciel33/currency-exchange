package business

import (
	"fmt"
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/models"
	"log"
	"strconv"
)

type ProductService struct {
	repository ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return ProductService{repo}
}

func (p ProductService) CreateNewProduct(data models.AddProduct) error {
	log.Println("ProductService::Create new product")
	product := AddProductToEntity(data)
	log.Println(fmt.Sprintf("Product: %v", product))

	if err := p.repository.Save(product); err != nil {
		return err
	}

	return nil
}

func (p ProductService) FindProductById(value string) (models.Product, error) {
	entity, err := p.repository.FindById(value)

	if err != nil {
		return models.Product{}, err
	}

	return addCurrencyExchange(entity), nil
}

func (p ProductService) ListAll() []models.Product {

	products := p.repository.ListAllProducts()
	productsWithExchange := make([]models.Product, 0)


	if products != nil {
		for _, product := range products {

			productsWithExchange = append(productsWithExchange, addCurrencyExchange(product))
		}
	}

	return productsWithExchange
}

func (p ProductService) UpdateProduct(data models.Product) error {
	log.Println("ProductService::Update")

	_, err := p.repository.FindById(data.Id.String())
	if err != nil {
		return err
	}

	product := UpdateProductToEntity(data)

	if err := p.repository.Update(product); err != nil {
		return err
	}

	return nil
}
func addCurrencyExchange(entity models.Product) models.Product {

	currencies, _ := GetCurrencies()

	if currencies != nil{
		for _, currency := range currencies {
			s, err := strconv.ParseFloat(currency.Value, 64)
			if err != nil {
				fmt.Print(err.Error())
			}
			exchange := entity.Price * s
			entity.Values =append(entity.Values, models.CurrencyExchange{Code: currency.Code, Value: exchange})
		}
	}

	return entity
}