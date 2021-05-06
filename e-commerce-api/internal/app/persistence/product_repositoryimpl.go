package persistence

import (
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/business"
	"github.com/jmaciel33/currency-exchange/e-commerce-api/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	INSERT   = "INSERT INTO products (product_id, name, price) VALUES ($1, $2, $3)"
	UPDATE   = "UPDATE products SET name = $1, price = $2 WHERE product_id = $3"
	DELETE   = "DELETE FROM products WHERE product_id = $1"
	FIND_ID  = "SELECT * FROM products WHERE product_id = $1"
	LIST_ALL = "SELECT * FROM products"
)

type ProductRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) business.ProductRepository {
	return ProductRepositoryImpl{db}
}

func (p ProductRepositoryImpl) Save(product models.Product) error {
	transaction := p.db.MustBegin()
	transaction.MustExec(INSERT, product.Id, product.Name, product.Price)
	if err := transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func (p ProductRepositoryImpl) Update(product models.Product) error {
	transaction := p.db.MustBegin()
	transaction.MustExec(UPDATE, product.Name, product.Price, product.Id)
	if err := transaction.Commit(); err != nil {
		return err
	}

	return nil

}

func (p ProductRepositoryImpl) FindById(value string) (models.Product, error) {
	var entity models.Product
	if err := p.db.Get(&entity, FIND_ID, value); err != nil {
		return models.Product{}, err
	}

	return entity, nil

}

func (p ProductRepositoryImpl) ListAllProducts() []models.Product {
	var products []models.Product
	if err := p.db.Select(&products, LIST_ALL); err != nil {
		log.Printf("[WARN] %q", err)
		return make([]models.Product, 0)
	}
	return products
}

func (p ProductRepositoryImpl) Delete(value string) error {
	transaction := p.db.MustBegin()
	transaction.MustExec(DELETE, value)
	if err := transaction.Commit(); err != nil {
		return err
	}
	return nil
}
